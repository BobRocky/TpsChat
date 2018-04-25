package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/richard-lyman/lithcrypt" //Ссылка алгоритма шифрования
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
	conn, err = net.Dial ("tcp", "lacale host:1313")

	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	for {
		sourse = Scan1()

		if (sourse == "exit" || sourse == "Exit")
		{
			os.Exit(1)
		}

		if n, err := conn.Write(encrypted); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

	}


}
	

