package handlers

import (
	"log"
	"net/http"

	"github.com/ManuEduardo/save-paste/src/views"
	"github.com/ManuEduardo/save-paste/src/views/components"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	homeProps := views.HomeProps{
		TagsItems: []components.TagItemProps{
			{TitleFilter: "prueba", IsAdded: false},
			{TitleFilter: "prueba 2", IsAdded: false},
			{TitleFilter: "prueba xd", IsAdded: false},
			{TitleFilter: "prueba owo", IsAdded: false},
			{TitleFilter: "prueba ñya", IsAdded: false},
		},
		SavesCards: []components.SaveCardProps{
			{Title: "Pruebita uwu",
				Content: "En teoria este deberia ser un super textaso que a saber lluego como voy a formatear ptmr",
				Tags: []components.TagItemProps{
					{TitleFilter: "prueba", IsAdded: false},
					{TitleFilter: "prueba 2", IsAdded: false},
				},
				DateCreated: "2022-02-02",
				IsVisible:   true,
			},
			{
				Content: "jfdslkfjlsakdf",
			},
		},
	}
	log.Println("Using Home Handler")
	views.Home(homeProps).Render(r.Context(), w)
}
