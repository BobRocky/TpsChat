package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/richard-lyman/lithcrypt"
)

var source string
var n int
var input []byte
var encrypted []byte
var proverka string
var otvet string

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func shifr(conn net.Conn) string {
	//ключ
	password := []byte("some")

	defer conn.Close()
	//Принятие слова
	input := make([]byte, (1024))
	n, err := conn.Read(input)
	if n == 0 || err != nil {
		fmt.Println("Read error:", err)
		os.Exit(1)
	}

	//Расшифруем
	original, decrypt_error := lithcrypt.Decrypt(password, input[0:n])
	if decrypt_error != nil {
		fmt.Println("Failed to decrypt:", decrypt_error)
		os.Exit(1)
	}

	source := string(original)
	//Слово в source
	return source
}

func main() {

	password := []byte("some")

	conn, err := net.Dial("tcp", "localhost:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	proverka := shifr(conn)
	if proverka == "errl" { //errl Ошибка логина
		fmt.Println("Логин введен не верно\nПовторите попытку: ")
	}

	for {
		fmt.Print("Password: ")
		source = Scan1()

		//************************Шифруем********************************************
		payload := []byte(source)
		encrypted, encrypt_error := lithcrypt.Encrypt(password, payload)
		if encrypt_error != nil {
			fmt.Println("Failed to encrypt:", encrypt_error)
			os.Exit(1)
		}
		//***************************************************************************

		//******************Отправка сообщения серверу******************
		if n, err := conn.Write(encrypted); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		proverka := shifr(conn)
		if proverka == "errp" { //errl Ошибка логина
			fmt.Println("Пароль введен не верно\nПовторите попытку: ")
		} else {
			break
		}

		for {
			source = Scan1()

			if source == "/help" {
				fmt.Println("/l - список челиков онлайн\n/w id написать челику")
			}
			//******************Шифруем**********************************************
			payload := []byte(source)
			encrypted, encrypt_error := lithcrypt.Encrypt(password, payload)
			if encrypt_error != nil {
				fmt.Println("Failed to encrypt:", encrypt_error)
				os.Exit(1)
			}
			//***************************************************************************

			//******************Отправка сообщения серверу******************
			if n, err := conn.Write(encrypted); n == 0 || err != nil {
				fmt.Println(err)
				return
			}

			//*****************************************************************
			fmt.Println("Что-то говорит тебе ")
			fmt.Println(otvet)

		}

	}
}
