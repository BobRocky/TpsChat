package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tsp", ":8080")

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

func handleConnection(conn net.Conn) {

	password := []byte("some")
}
