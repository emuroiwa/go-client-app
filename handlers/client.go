package handlers

import (
	"html/template"
	"net/http"
	"net/mail"
	"strconv"

	"go-client-app/models"
)

type ClientHandler struct {
	store     models.ClientStore
	templates *template.Template
}

func NewClientHandler(store models.ClientStore, tmpl *template.Template) *ClientHandler {
	return &ClientHandler{store, tmpl}
}

func (h *ClientHandler) List(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "index.html", h.store.All())
}

func (h *ClientHandler) ShowForm(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "form.html", nil)
}

func (h *ClientHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	if _, err := mail.ParseAddress(email); err != nil {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	h.store.Create(models.Client{Name: name, Email: email})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *ClientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	h.store.Delete(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
