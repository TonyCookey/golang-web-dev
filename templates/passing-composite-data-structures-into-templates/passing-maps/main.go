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
	footballers := map[string]string{
		"Premier League/ Seria A/Portugal":    "Ronaldo",
		"Ligue 1/ La Liga Satander/Argentina": "Messi",
		"Bundesliga/Poland":                   "Lewandowski",
		"Ligue 1/ Brazil":                     "Neymar",
		"Ligue 1/ France":                     "Mbappe",
		"Bundesliga/Norway":                   "Haaland",
	}
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", footballers)
	if err != nil {
		log.Fatalln(err)
	}
}
