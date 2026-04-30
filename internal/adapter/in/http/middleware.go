package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func UseDefaultMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}