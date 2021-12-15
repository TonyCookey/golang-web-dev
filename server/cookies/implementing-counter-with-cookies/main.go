package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("cookie-counter")
	fmt.Println(c, err)

	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "cookie-counter",
			Value: strconv.Itoa(1),
			Path:  "/",
		})
		return
	}

	counter, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	counter++
	c.Value = strconv.Itoa(counter)
	http.SetCookie(w, c)
	//io.WriteString(w, c.Value)
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
