package handlers

import (
	"log"
	"net/http"

	"github.com/ManuEduardo/save-paste/src/views"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func HandleBase(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Base Handler")
	views.BasePage().Render(r.Context(), w)
}
