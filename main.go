package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		// Handle error
		return
	}
	t.Execute(w, nil) // No data needed for this simple template
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch email me at nurfaiz.foat[at]gmail.com</a>.")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Memuat di pelayan :8080...")
	http.ListenAndServe(":8080", router)
}
