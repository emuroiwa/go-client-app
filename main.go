package main

import (
	"html/template"
	"net/http"

	"go-client-app/handlers"
	"go-client-app/models"

	"log"

	"github.com/go-chi/chi/v5"
)

func main() {
	tmpl := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/index.html",
		"templates/form.html",
	))

	store := models.NewInMemoryClientStore()
	clientHandler := handlers.NewClientHandler(store, tmpl)

	r := chi.NewRouter()
	r.Get("/", clientHandler.List)
	r.Get("/new", clientHandler.ShowForm)
	r.Post("/create", clientHandler.Create)
	r.Get("/delete", clientHandler.Delete)

	log.Println("Server running")

	http.ListenAndServe(":3001", r)
}
