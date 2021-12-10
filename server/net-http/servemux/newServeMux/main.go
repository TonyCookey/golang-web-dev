package main

import (
	"io"
	"net/http"
)

type homepage string

func (h homepage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Homepage")
}

type aboutpage string

func (a aboutpage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "About us page")
}

func main() {
	var home homepage
	var about aboutpage

	mux := http.NewServeMux()
	mux.Handle("/home/", home)
	mux.Handle("/about", about)

	http.ListenAndServe(":8080", mux)
}
