package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("../views/login.html"))
	log.Println("Using Base Handler")
	template.Execute(w, nil)
}
