package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog", indexHandler)
	http.HandleFunc("/me", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hello world")
}
func meHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Anthony")
}
func dogHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hello dog")
}
