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
	award := struct {
		AwardName     string
		Footballer    footballer
		EventLocation string
	}{
		AwardName:     "Ballon D'Or 2021",
		Footballer:    bestFootballer,
		EventLocation: "Paris, France",
	}
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", award)
	if err != nil {
		log.Fatalln(err)
	}
}
