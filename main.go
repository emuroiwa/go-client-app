package main

import (
	"html/template"
	"log"
	"net/http"

	"go-client-app/handlers"
	"go-client-app/models"
	"go-client-app/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	store := models.NewInMemoryClientStore()
	service := services.NewClientService(store)
	clientHandler := handlers.NewClientHandler(service, tmpl)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", clientHandler.List)
	r.Post("/create", clientHandler.Create)
	r.Get("/delete", clientHandler.Delete)

	log.Println("Server started on :3002")
	if err := http.ListenAndServe(":3002", r); err != nil {
		log.Fatal("Server error:", err)
	}
}
