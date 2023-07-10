package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AlexTLDR/WebDev/views"
	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", "err")
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// tpl, err := template.ParseFiles("templates/home.gohtml") -> this is only working in unix systems
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	// tpl, err := template.ParseFiles("templates/contact.gohtml") -> this is only working in unix systems
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", r)
}
