package main

import (
	"bufio"
	"fmt"
	"os"
	//Ссылка алгоритма шифрования
)

/*
func main() {
	var x [5]int //Создаем массив с длинной в 5 делений
	x[4] = 100   // 4 символу даем знаение 100 (Счет идет с 0)
	fmt.Println(x)
}

func main() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(x)
	fmt.Println(total)
	fmt.Println(total / 5)
	fmt.Println(len(x))
}*/
/*func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}
func main() {
	slice1 := []int{1, 2, 3} //[1 2 3]
	slice2 := make([]int, 2) //[0 0]
	//copy(slice2, slice1) // [1 2 3] [1 2]
	copy(slice1, slice2) // [0 0 3] [0 0]
	fmt.Println(slice1, slice2)
}

func main() {
	//var x float64
	//var arr float64
	arr := [5]float64{1, 2, 3, 4, 5}
	//x := arr[0:5]
	x := arr[2:5]
	fmt.Println(arr, x)
}*/

/*func main(){
	var client string
	fmt.Scan(&client)

	client = make

}*/
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
func main() {
	var client string
	fmt.Print("Пиши уебок: ")
	client = Scan1()
	slice1 := []byte(client)
	slice2 := make([]byte, 2)
	copy(slice2, slice1)
	/*fmt.Println(slice1)         //[208 159 209 128 208 184 208 178 208 181 209 130]
	fmt.Println(string(slice1)) //Привет
	fmt.Println(slice2)  */ //[208 159]
	fmt.Print("Вот твой первый высер: ")
	fmt.Println(string(slice2)) //П

}
