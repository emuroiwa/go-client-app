package utils

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl *template.Template, name string, data any) {
	w.Header().Set("Content-Type", "text/html")
	err := tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println("Template rendering error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
