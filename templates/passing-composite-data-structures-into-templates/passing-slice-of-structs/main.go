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
	ronaldo := footballer{
		Name:     "Cristiano Ronaldo",
		League:   "Premier League",
		Country:  "Portugal",
		Age:      35,
		Position: "Striker",
	}
	messi := footballer{
		Name:     "Lionel Messi",
		League:   "Ligue 1",
		Country:  "Argentina",
		Age:      33,
		Position: "Striker",
	}
	lewandowski := footballer{
		Name:     "Robert Lewandowski",
		League:   "Bundesliga",
		Country:  "Poland",
		Age:      32,
		Position: "Striker",
	}
	footballers := []footballer{ronaldo, messi, lewandowski}
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", footballers)
	if err != nil {
		log.Fatalln(err)
	}
}
