package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", coin)
	http.HandleFunc("/coin.jpg", coinPic)
	http.ListenAndServe(":8080", nil)
}

func coin(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="coin.jpg">	`)
}

func coinPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "coin.jpg")
}
