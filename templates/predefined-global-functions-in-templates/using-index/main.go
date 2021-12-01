package main

import (
	"log"
	"os"
	"text/template"
)

var templateFile *template.Template

func init() {
	templateFile = template.Must(template.ParseFiles("main.gohtml"))
}
func main() {
	footballers := []string{"Ronaldo", "Messi", "Lewandowski", "Neymar", "Mbappe", "Haaland"}
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", footballers)
	if err != nil {
		log.Fatalln(err)
	}
}
