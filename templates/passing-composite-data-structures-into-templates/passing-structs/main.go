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

type footballer struct {
	Name     string
	League   string
	Country  string
	Age      int
	Position string
}

func main() {
	bestFootballer := footballer{
		Name:     "Cristiano Ronaldo",
		League:   "Premier League",
		Country:  "Portugal",
		Age:      35,
		Position: "Striker",
	}
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", bestFootballer)
	if err != nil {
		log.Fatalln(err)
	}
}
