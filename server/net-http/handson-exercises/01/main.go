package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dog", dogHandler)
	http.HandleFunc("/me", meHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(res http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(res, "hello world")
	if err != nil {
		log.Fatalln(err)
	}

}
func meHandler(res http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(res, "Anthony")
	if err != nil {
		log.Fatalln(err)
	}

}
func dogHandler(res http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(res, "hello dog")
	if err != nil {
		log.Fatalln(err)
	}
}
