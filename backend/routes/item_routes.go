package routes

import (
	"github.com/go-chi/chi/v5"
	"antis/backend/handlers"
)

func SetupItemRoute(r chi.Router, itemHandler *handlers.ItemHandler){
	r.Route("/items", func(r chi.Router) {
		r.Get("/", itemHandler.GetAllItems)
	})
}