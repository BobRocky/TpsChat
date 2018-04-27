package main

import (
	"fmt"
	"net"
	"os"

	"github.com/richard-lyman/lithcrypt"
)

var source string
var n int
var input []byte
var encrypted []byte

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
func scan2(input []byte) {
	password := []byte("some")
	original, decrypt_error := lithcrypt.Decrypt(password, input[0:n])
	if decrypt_error != nil {
		fmt.Println("Failed to decrypt:", decrypt_error)
		os.Exit(1)
	}
	source = string(original)
	return
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

//Обработка подключения
func handleConnection(conn net.Conn) {
	//Ключ

	for {
		//Считвыем полученные данные
		input := make([]byte, (1024))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		//Расшифровываем
		scan2(input)
		fmt.Print(input)
		if source != "Kirill" {
			fmt.Println(scan2, source)
		}
	}
}
