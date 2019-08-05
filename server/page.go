package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (templ templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templ.once.Do(func() {
		templ.templ = template.Must(template.ParseFiles(filepath.Join("templates", templ.filename)))
	})
	templ.templ.Execute(w, nil)
}

func main() {
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
