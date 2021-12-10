package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", dog)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func dog(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `<img src="/toby.jpg">`)
	if err != nil {
		log.Fatalln(err)
	}
}
