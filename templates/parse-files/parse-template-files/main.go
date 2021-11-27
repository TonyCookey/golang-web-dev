package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	templateFile, err := template.ParseFiles("main.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = templateFile.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
