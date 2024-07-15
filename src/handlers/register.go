package handlers

import (
	"log"
	"net/http"

	"github.com/ManuEduardo/save-paste/src/views"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Register Handler")
	views.Register().Render(r.Context(), w)
}
