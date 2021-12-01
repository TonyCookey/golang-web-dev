package main

import (
	"log"
	"os"
	"text/template"
)

var templateFile *template.Template

func init() {
	templateFile = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	data := struct {
		Header string
		Body   string
		Footer string
	}{
		Header: "This is my Header",
		Body:   "Christ is the Head of the Church, and Upon that head, I will build my Church, and the gates of hell shall not prevail",
		Footer: "This is my Footer",
	}
	err := templateFile.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
