package main

import (
	"fmt"
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

}
