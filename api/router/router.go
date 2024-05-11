package router

import (
	"github.com/go-chi/chi/v5"

	"GoLibraryAPI/api/resource/book"
	"GoLibraryAPI/api/resource/health"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/status", health.Read)

	r.Route("/v1", func(r chi.Router) {
		bookAPI := &book.API{}
		r.Get("/books", bookAPI.List)
		r.Post("/books", bookAPI.Create)
		r.Get("/books/{id}", bookAPI.Read)
		r.Put("/books/{id}", bookAPI.Update)
		r.Delete("/books/{id}", bookAPI.Delete)
	})

	return r
}
