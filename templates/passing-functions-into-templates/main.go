package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var templateFile *template.Template

var funcmap = template.FuncMap{
	"stringToUpper": strings.ToUpper,
	"toLowerCase":   toLowerCase,
}

func init() {
	templateFile = template.Must(template.New("").Funcs(funcmap).ParseFiles("main.gohtml"))
}
func toLowerCase(s string) string {
	return strings.ToLower(s)
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
		Position: "striker",
	}
	messi := footballer{
		Name:     "Lionel Messi",
		League:   "Ligue 1",
		Country:  "Argentina",
		Age:      33,
		Position: "striker",
	}
	lewandowski := footballer{
		Name:     "Robert Lewandowski",
		League:   "Bundesliga",
		Country:  "Poland",
		Age:      32,
		Position: "striker",
	}
	footballers := []footballer{ronaldo, messi, lewandowski}
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", footballers)
	if err != nil {
		log.Fatalln(err)
	}
}
