/*package main

import (
	"os"
	"fmt"
	"net"
)

func main()
{
	listener, err := net.Listen("tsp", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer istener.Close()
	fmt.Println("Startig server...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
	}
}