package handlers

import (
	"log"
	"net/http"

	"github.com/ManuEduardo/save-paste/src/views"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Base Handler")
	views.Login("Manu").Render(r.Context(), w)
}
