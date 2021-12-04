package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	err := connection.SetDeadline(time.Now().Add(100 * time.Second))
	if err != nil {
		log.Fatalln("Connection Timed Out after 100 Seconds")
	}
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		requestLine := scanner.Text()
		fmt.Println(requestLine)
		_, err = fmt.Fprintln(connection, "Type something...")
		if err != nil {
			log.Fatalln(err)
		}

	}
	defer connection.Close()

	fmt.Println("Code Ended")
}
