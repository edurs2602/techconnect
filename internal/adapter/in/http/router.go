package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(userHandler *UserHandler) http.Handler {
	r := chi.NewRouter()

	//r.Use(chiMiddleware.Logger)

	r.Post("/auth/register", userHandler.Register)

	return r
}

func respond(w http.ResponseWriter, status int, data any) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondErr(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
