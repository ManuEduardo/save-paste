package handlers

import (
	"log"
	"net/http"

	"github.com/ManuEduardo/save-paste/src/views"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Using Home Handler")
	views.Home().Render(r.Context(), w)
}
