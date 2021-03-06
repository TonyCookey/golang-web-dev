package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", coin)
	http.HandleFunc("/coin.jpg", coinPic)
	http.ListenAndServe(":8080", nil)
}

func coin(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="/coin.jpg">
	`)
}

func coinPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("coin.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
