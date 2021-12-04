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
		go handle(connection)
	}
}
func handle(connection net.Conn) {
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		requestLine := scanner.Text()
		fmt.Println(requestLine)
	}
	defer connection.Close()
	fmt.Println("Chat Ended")
}
