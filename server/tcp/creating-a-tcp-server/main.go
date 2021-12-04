package main

import (
	"fmt"
	"io"
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
		io.WriteString(connection, "\n Hello World from a Server \n")
		fmt.Fprintf(connection, "%v", time.Now())
		fmt.Fprintln(connection, "How is your day going?")

		connection.Close()
	}
}
