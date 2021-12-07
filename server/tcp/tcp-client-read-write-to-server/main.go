package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close()
	output, err := ioutil.ReadAll(connection)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(output))
	fmt.Fprintln(connection, "Testing Dialing in to connections")

}
