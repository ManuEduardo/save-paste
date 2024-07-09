package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func HandleBase(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("src/views/basePage.html"))
	log.Println("Using Base Handler")
	template.Execute(w, nil)
}
