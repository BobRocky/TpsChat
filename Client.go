package main

import (
	"fmt"
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
}

/*func main(){
	var client string
	fmt.Scan(&client)

	client = make

}
func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
}
<<<<<<< HEAD
	func main() {
	var client string
	fmt.Print("Пиши уебок: ")
	client = Scan1()
	slice1 := []byte(client)
	slice2 := make([]byte, 2)
	copy(slice2, slice1)
	/*fmt.Println(slice1)         	//[208 159 209 128 208 184 208 178 208 181 209 130]
	fmt.Println(string(slice1))		 //Привет
	fmt.Println(slice2)				 //[208 159]
	fmt.Print("Вот твой первый высер: ")
	fmt.Println(string(slice2)) 	//П

}*/
func main() {
	arr := []float64{1, 2, 3, 4, 5}
	x := arr[0:5]
	fmt.Println(arr, x)
=======

//*******************************************************************

func main() {
	//*****************************Key*************************************
	password := []byte("some")
	//****************************************************

	//***************************Подключение\Проверка*************************
	conn, err := net.Dial("tcp", "localhost:4545")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()
	//********************************************************************
	for {
		fmt.Print("Login: ")
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
>>>>>>> bac50106d1f7a2e1e3be46031534e6d2769fb46a
}
