package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/richard-lyman/lithcrypt" //Ссылка алгоритма шифрования
)

var source string
var proverka string

//*************Функция ввода слова*********************

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

//*******************************************************************

func main() {
	//*****************************Key*************************************
	password := []byte("some")
	//****************************************************

	//***************************Подключение\Проверка*************************
	conn, err := net.Dial("tcp", "25.30.186.173:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	//********************************************************************
	for {
		fmt.Println("login: ")
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
		//*****************************************************************

		//*********************Полученный ответ*****************************
		buff := make([]byte, 8192)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		//***************************************************************************

		//*************************Расшифруем***************************
		original, decrypt_error := lithcrypt.Decrypt(password, buff[0:n])
		if decrypt_error != nil {
			fmt.Println("Failed to decrypt:", decrypt_error)
			os.Exit(1)
		}
		//**************************************************************

		proverka = string(original)
		if proverka == "errl" { //errl Ошибка логина
			fmt.Println("Логин введен не верно\nПовторите попытку: ")
		} else {
			break
		}

	}

	for {
		fmt.Println("password: ")
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
		//*****************************************************************

		//*********************Полученный ответ*****************************
		buff := make([]byte, 8192)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		//***************************************************************************

		//*************************Расшифруем***************************
		original, decrypt_error := lithcrypt.Decrypt(password, buff[0:n])
		if decrypt_error != nil {
			fmt.Println("Failed to decrypt:", decrypt_error)
			os.Exit(1)
		}
		//**************************************************************

		proverka = string(original)
		if proverka == "errp" { //errl Ошибка логина
			fmt.Println("Пароль введен не верно\nПовторите попытку: ")
		} else {
			break
		}

		for {
			source = Scan1()
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

			//*********************Полученный ответ*****************************
			buff := make([]byte, 8192)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			//***************************************************************************

			//*************************Расшифруем***************************
			original, decrypt_error := lithcrypt.Decrypt(password, buff[0:n])
			if decrypt_error != nil {
				fmt.Println("Failed to decrypt:", decrypt_error)
				os.Exit(1)
			}
			//**************************************************************
			fmt.Println("Что-то говорит тебе ")
			fmt.Println(string(original))

		}

	}
}
