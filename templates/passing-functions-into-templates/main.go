package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var templateFile *template.Template

var funcmap = template.FuncMap{
	"stringToUpper": strings.ToUpper,
	"toLowerCase":   toLowerCase,
	"formatTime":    formatTime,
	"getYear":       getYear,
}

func init() {
	templateFile = template.Must(template.New("").Funcs(funcmap).ParseFiles("main.gohtml"))
}
func formatTime(t time.Time) string {
	return t.Format("01-02-2006")
}

func getYear(t time.Time) string {
	return t.Format("2006")

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
	ceremonyDetails := struct {
		AwardName     string
		Footballers   []footballer
		EventLocation string
		EventTime     time.Time
	}{
		AwardName:     "Ballon D'Or",
		Footballers:   footballers,
		EventLocation: "Paris, France",
		EventTime:     time.Now(),
	}
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", ceremonyDetails)
	if err != nil {
		log.Fatalln(err)
	}
}
