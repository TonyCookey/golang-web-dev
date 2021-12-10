package main

import (
	"io"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Home page")
}

func about(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "About us Page")
}

func main() {

	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8080", nil)
}
