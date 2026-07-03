package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type TokenValidator interface {
	Validate(token string) (string, error)
}

func NewRouter(userHandler *UserHandler, postHandler *PostHandler, jwtSvc TokenValidator) http.Handler {
	r := chi.NewRouter()

	UseDefaultMiddlewares(r)
	r.Use(NewRateLimiter(3, 10*time.Second))

	r.Get("/health", Health)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
		r.Post("/login", userHandler.Login)
		r.Post("/refresh", userHandler.Refresh)
	})

	r.Route("/user", func(r chi.Router) {
		r.Get("/{username}", userHandler.GetUser)
		//	r.Get("/{username}/posts", postHandler.ListByUser)
		r.Patch("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})

	r.Group(func(r chi.Router) {
		r.Use(Auth(jwtSvc))

		r.Patch("/users/me", userHandler.UpdateUser)
		r.Delete("/users/me", userHandler.DeleteUser)

		r.Route("/posts", func(r chi.Router) {
			r.Get("/", postHandler.List)
			r.Post("/", postHandler.Create)
			r.Get("/{id}", postHandler.GetByID)
			r.Delete("/{id}", postHandler.Delete)
			r.Post("/{id}/comments", postHandler.AddComment)
			r.Delete("/{id}/comments/{commentId}", postHandler.DeleteComment)
		})
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
