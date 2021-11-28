package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	templateFile, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}
	err = templateFile.Execute(os.Stdout, nil)
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
