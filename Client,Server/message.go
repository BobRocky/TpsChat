package chat

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (self *Message) String() string {
	return self.Author + " говорит: " + self.Body
}

/*
func shifr(conn net.Conn) string {
	//ключ
	password := []byte("some")

	defer conn.Close()
	//Принятие слова
	input := make([]byte, (1024))
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

	source := string(original)
	//Слово в source
	return source
}