package handlers

import (
	"log"
	"net/http"

	"github.com/ManuEduardo/save-paste/src/views"
	"github.com/ManuEduardo/save-paste/src/views/components"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	homeProps := views.HomeProps{
		Items: []components.TagItemProps{
			{TitleFilter: "Prueba"},
			{TitleFilter: "Prueba xdd"},
		},
	}
	log.Println("Using Home Handler")
	views.Home(homeProps).Render(r.Context(), w)
}
