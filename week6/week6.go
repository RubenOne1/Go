package main

import (
	"net/http"
	"html/template"
)

type person struct {
	First string
	Last string
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
  http.HandleFunc("/info", info)
  http.HandleFunc("/media", media)
  http.HandleFunc("/contact", contact)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "Hello World!")
}

func about(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", 42)
}
func info(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		"James",
		"Bond",
	}

	tpl.ExecuteTemplate(w, "info.gohtml", p1)
}

func media(w http.ResponseWriter, r *http.Request) {
	xi := []int{3,5,7,9,17,749}
	tpl.ExecuteTemplate(w, "media.gohtml", xi)
}

func contact(w http.ResponseWriter, r *http.Request) {
	m := map[string]int{
		"James":32,
		"Moneypenny":24,
	}
	tpl.ExecuteTemplate(w, "contact.gohtml", m)
}
