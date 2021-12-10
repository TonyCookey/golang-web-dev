package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var templateFile *template.Template

func init() {
	templateFile = template.Must(template.ParseGlob("templates/*"))
}
func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/profile", profileHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(res http.ResponseWriter, _ *http.Request) {
	err := templateFile.ExecuteTemplate(res, "index.gohtml", "tony cookey")
	if err != nil {
		log.Fatalln(err)
	}
}
func aboutHandler(res http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(res, "Software Engineer")
	if err != nil {
		log.Fatalln(err)
	}
}
func profileHandler(res http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(res, "Tony Cookey")
	if err != nil {
		log.Fatalln(err)
	}
}
