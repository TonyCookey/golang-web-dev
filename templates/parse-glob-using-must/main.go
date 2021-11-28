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
	err := templateFile.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = templateFile.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = templateFile.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
