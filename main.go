package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	days := []int{1, 2, 3, 4, 5, 6, 7}
	t.Execute(w, days)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}

	http.HandleFunc("/", index)
	server.ListenAndServe()
}
