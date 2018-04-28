package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/richard-lyman/lithcrypt"
)

var source string
var errl string
var errp string
var noerr string
var a string

/*
var source string
var n int
var input []byte
var encrypted []byte
var original []byte
var conn net.Conn

//Функция шифрования
func scan1(encrypted []byte) {

	password := []byte("some")
	payload := []byte(source)
	encrypted, encrypt_error := lithcrypt.Encrypt(password, payload)
	if encrypt_error != nil {
		fmt.Println("Failed to encrypt:", encrypt_error)
		os.Exit(1)
	}
	return
}

//функция расшифровки
func scan2() {

	input := make([]byte, (8192))
	n, err := conn.Read(input)
	if n == 0 || err != nil {
		fmt.Println("Read error:", err)
		os.Exit(1)
	}

	password := []byte("some")
	original, decrypt_error := lithcrypt.Decrypt(password, input[0:n])
	if decrypt_error != nil {
		fmt.Println("Failed to decrypt:", decrypt_error)
		os.Exit(1)
	}
	source = string(original)
	return
}*/

func Scan2() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Startig server...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

func shifr(conn net.Conn) string {
	//ключ
	password := []byte("some")

	//Принятие слова
	input := make([]byte, (1024 * 8))
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

	//source := string(original)
	//Слово в source
	return string(original)
}

func shifr1(conn net.Conn) string {

	password := []byte("some")

	payload := []byte(source)
	encrypted, encrypt_error := lithcrypt.Encrypt(password, payload)
	if encrypt_error != nil {
		fmt.Println("Failed to encrypt:", encrypt_error)
		os.Exit(1)
	}
	if n, err := conn.Write(encrypted); n == 0 || err != nil {
		fmt.Print("Client: ")
		fmt.Println(err)
		os.Exit(1)
	}

	return source
}

//Обработка подключения
func handleConnection(conn net.Conn) {

	defer conn.Close()

	for {

		a = shifr(conn)
		fmt.Println(a)

		if a != "Kirill" {

			errl = "errl"
			source = errl
			source := shifr1(conn)
			fmt.Println(source)
			continue

		} else {

			noerr = "noerr"
			source = noerr
			source := shifr1(conn)
			fmt.Println(source)
			break
		}
	}
	for {
		// считываем полученные в запросе данные
		a = shifr(conn)
		fmt.Println(a)

		if a != "Kirill" {

			errp := "errp"
			source = errp
			source := shifr1(conn)
			fmt.Println(source)
			continue

		} else {

			noerr = "noerr"
			source = noerr
			source := shifr1(conn)
			fmt.Println(source)
			break

		}
	}
}
