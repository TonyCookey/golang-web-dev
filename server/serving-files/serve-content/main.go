package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", coin)
	http.HandleFunc("/coin.jpg", coinPic)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func coin(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/coin.jpg">`)
}

func coinPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("coin.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
