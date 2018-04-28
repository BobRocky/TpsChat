/*package main

import (
	"bufio"
	"fmt"
	"os"
)
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
var q = Scan1() //Объявляем переменную вне функции

func proverka() string {
	arr := []byte(q) //из переменной q делаем массив arr
	y := arr[0:2]    //из arr вырезаем символы от 0 до 2 в переменную y
	return string(y) //Функцие proverka даем значение y в расширении string
}

func proverka2() string {
	arr := []byte(q)     //из переменной q делаем массив arr
	x := arr[5:len(arr)] //переменной даем символы от arr от 5 символа до длины arr
	return string(x)     //Функцие proverka2 даем значение y в расширении string
}
func proverka3() string {
	arr := []byte(q) //из переменной q делаем массив arr
	y := arr[3:4]    //переменной даем символы от arr от 3 до 4 символа
	return string(y) //Функцие proverka даем значение y в расширении string
}

func main() {
	var comand string
	var slovo string
	var id string
	comand = proverka()
	if comand == "/w" {
		id = proverka3()
		slovo = proverka2()
		fmt.Print("Пользвателю ", id, ": ", slovo)
	}

}*/

package main

func main() {

}
