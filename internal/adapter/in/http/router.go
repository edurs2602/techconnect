package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(userHandler *UserHandler, postHandler *PostHandler) http.Handler {
	r := chi.NewRouter()

	UseDefaultMiddlewares(r)

	r.Get("/health", Health)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
	})

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", postHandler.List)
		r.Post("/", postHandler.Create)
		r.Get("/{id}", postHandler.GetByID)
	})

	return r
}

func Health(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "TechConnect API rodando",
	})
}

func respond(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondErr(w http.ResponseWriter, status int, msg string) {
	respond(w, status, map[string]string{"error": msg})
}