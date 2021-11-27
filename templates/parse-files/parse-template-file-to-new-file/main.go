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
	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = templateFile.Execute(newFile, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
