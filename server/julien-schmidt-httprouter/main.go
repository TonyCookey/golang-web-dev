package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.HandlerFunc("GET", "/testing", test)
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)
	mux.POST("/blog/:category/:article", blogWrite)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
func test(_ http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Println("Hello, 世界")
	if err != nil {
		log.Fatalln(err)
	}
}
func user(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	_, err := fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
	if err != nil {
		log.Fatalln(err)
	}
}

func blogRead(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	_, err := fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
	if err != nil {
		log.Fatalln(err)
	}
}

func blogWrite(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	_, err := fmt.Fprintf(w, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintf(w, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
