package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/caution.jpg", Nuke)
	http.HandleFunc("/cat.jpg", Cat)
	http.HandleFunc("/dragon.jpg", Dragon)
	http.HandleFunc("/heart.jpg", Heart)
	http.HandleFunc("/lighter.jpg", Lighter)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="caution.jpg">	`)
}

func Nuke(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "more/caution.jpg")
}
func Cat(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "more/cat.jpg")
}
func Dragon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "more/dragon.jpg")
}
func Heart(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "more/heart.jpg")
}
func Lighter(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "more/lighter.jpg")
}
