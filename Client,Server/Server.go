package main

import (
	"fmt"
	"net"
	"os"

	"github.com/richard-lyman/lithcrypt"
)

var dict = map[string]string{
	"fatty":           "322",
	"kista":           "qwertyasd",
	"old":             "piror228",
	"arsen":           "senya",
	"sutulaya_sobaka": "anya",
}

var source string
var errl string
var errp string
var noerr string
var a string
var q string
var b string

//Функция работы сервреа
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

//Функция принятия данных и рашифровки
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

//Функция шифровки и отправки
func shifr1(conn net.Conn) string {

	password := []byte("some")

	//Шифруем
	payload := []byte(source)
	encrypted, encrypt_error := lithcrypt.Encrypt(password, payload)
	if encrypt_error != nil {
		fmt.Println("Failed to encrypt:", encrypt_error)
		os.Exit(1)
	}

	//Отправляем
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

	//Цикл проверки логина
	for {

		//принимаем и расшифровываем логин
		a = shifr(conn)

		//Проверяем существует такой логин или нет
		login := string(a)
		target, a := dict[login]

		//Проверям правильность лониа
		if a == false {

			//Если логин не правильный отправляем клиенту ошибку
			errl = "errl"
			source = errl
			//Шифруем и отраляем ошибку
			source := shifr1(conn)
			fmt.Println("source1: ", source)
			//Продолжаем цикл
			continue

		} else {

			//Если логин правильный отправляем клиенту то что ошибок нету
			noerr = "noerr"
			source = noerr
			//Шифруем и отправляем то что ошибки нет
			source := shifr1(conn)
			fmt.Println("source2: ", source)
			//Закнчиваем цикл
			for {

				fmt.Println("target: ", target)
				//Принимаем и рашифровываем пароль
				b = shifr(conn)
				fmt.Println(b)

				//Проверяем правильность пароля
				if ne b target{

					//Если пароль не пральный отправляем ошибку
					errp := "errp"
					source = errp
					//Шифруем и отраляем ошибку
					source := shifr1(conn)
					fmt.Println(source)
					//Продолжаем цикл
					continue

				} else {

					noerr = "noerr"
					source = noerr
					//Шифруем и отправляем то что ошибки нет
					source := shifr1(conn)
					fmt.Println(source)
					//Закнчиваем цикл
					break

				}
			}
			break
		}
	}
	//Цикл проверки пароля

	for {

	}
}
