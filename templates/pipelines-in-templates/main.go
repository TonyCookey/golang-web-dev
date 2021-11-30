package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var templateFile *template.Template

var funcmap = template.FuncMap{
	"square":     square,
	"squareRoot": squareRoot,
	"triple":     triple,
}

func init() {
	templateFile = template.Must(template.New("").Funcs(funcmap).ParseFiles("main.gohtml"))
}
func square(x float64) float64 {
	return x * x
}
func squareRoot(x float64) float64 {
	return math.Sqrt(x)
}
func triple(x float64) float64 {
	return x * x * x
}

func main() {
	err := templateFile.ExecuteTemplate(os.Stdout, "main.gohtml", float64(9))
	if err != nil {
		log.Fatalln(err)
	}
}
