package main

import (
	"log"
	"os"
	"text/template"
)

var templateFile *template.Template

func init() {
	templateFile = template.Must(template.ParseFiles("one.gohtml"))
}
func main() {
	err := templateFile.ExecuteTemplate(os.Stdout, "one.gohtml", 7)
	if err != nil {
		log.Fatalln(err)
	}
}
