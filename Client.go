package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	//Ссылка алгоритма шифрования
)

var sourse string //Создания слова вводимое с клавиатуры

//Функция позволяющая засунуть в переменную более 1 слова.
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}

func main() {
	//Подключение к серверу
	conn, err = net.Dial("tcp", "lacale host:1313")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("login: ")
	var sourse = Scan1()

	//Отправка сообщение серверу
	if n, err = conn.Write(sourse); n == 0 || err != nil {
		fmt.Println(err)
		return
	}

	for buff != "errl" {
		buff := make([]byte, 8192)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		if buff == "errl" {
			fmt.Println("Логин введен неверно/nВведите логин заного: ")
		}
	}

	for buff != "errp" {
		buff := make([]byte, 8192)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		if buff == "errp" {
			fmt.Println("Пароль введен неверно/nВведите пароль заного: ")
		}
	}

	for {
		sourse = Scan1()

		if sourse == "exit" {
			os.Exit(1)
		}
		for {
			sourse = Scan1()

			if sourse == "Exit" {
				os.Exit(1)
			}

			for {
				//Отправка сообщение серверу
				if n, err := conn.Write(sourse); n == 0 || err != nil {
					fmt.Println(err)
					return
				}
			}

		}

	}
}
