package routes

import (
	"antis/backend/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(
	itemHandler *handlers.ItemHandler,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)

	r.Route("/api", func(r chi.Router) {
		 SetupItemRoute(r, itemHandler)
	})

	return r
}