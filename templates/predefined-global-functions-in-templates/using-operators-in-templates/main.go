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
	scores := struct {
		FirstScore  int
		SecondScore int
	}{
		FirstScore:  12,
		SecondScore: 8,
	}
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", scores)
	if err != nil {
		log.Fatalln(err)
	}
}
