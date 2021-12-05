package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		_, err = fmt.Fprintln(connection, "Group Chat Started")
		if err != nil {
			log.Fatalln(err)
		}

		go handle(connection)
	}
}
func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		requestLine := scanner.Text()
		fmt.Println(requestLine)
		_, err := fmt.Fprintln(connection, "type something...")
		if err != nil {
			log.Fatalln(err)
		}

	}
	defer connection.Close()

	fmt.Println("Code Ended")
}
