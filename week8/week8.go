package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
  http.HandleFunc("/friday", friday)
  http.HandleFunc("/saturday", saturday)
  http.HandleFunc("/sunday", sunday)
  http.HandleFunc("/monday", monday)
  http.HandleFunc("/tuesday", tuesday)
  http.HandleFunc("/wednesday", wednesday)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
func friday(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "friday.gohtml", nil)
}
func saturday(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "saturday.gohtml", nil)
}
func sunday(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "sunday.gohtml", nil)
}
func monday(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "monday.gohtml", nil)
}
func tuesday(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "tuesday.gohtml", nil)
}
func wednesday(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "wednesday.gohtml", nil)
}
