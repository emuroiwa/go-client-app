package handlers

import (
	"html/template"
	"log"
	"net/http"

	"go-client-app/services"
)

type ClientHandler struct {
	service   *services.ClientService
	templates *template.Template
}

func NewClientHandler(service *services.ClientService, tmpl *template.Template) *ClientHandler {
	return &ClientHandler{service, tmpl}
}

func (h *ClientHandler) List(w http.ResponseWriter, r *http.Request) {
	clients := h.service.ListClients()

	w.Header().Set("Content-Type", "text/html")
	if err := h.templates.ExecuteTemplate(w, "index.html", clients); err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func (h *ClientHandler) ShowForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := h.templates.ExecuteTemplate(w, "form.html", nil); err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if err := h.service.CreateClient(name, email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := h.service.DeleteClientByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
