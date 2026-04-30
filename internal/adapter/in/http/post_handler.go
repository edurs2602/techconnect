package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	mu     sync.Mutex
	posts  []PostResponse
	nextID int
}

type CreatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type PostResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		posts:  []PostResponse{},
		nextID: 1,
	}
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	respond(w, http.StatusOK, h.posts)
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input CreatePostInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondErr(w, http.StatusBadRequest, "payload inválido")
		return
	}

	if input.Title == "" {
		respondErr(w, http.StatusBadRequest, "título obrigatório")
		return
	}

	if input.Content == "" {
		respondErr(w, http.StatusBadRequest, "conteúdo obrigatório")
		return
	}

	if input.Author == "" {
		input.Author = "anônimo"
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	post := PostResponse{
		ID:        h.nextID,
		Title:     input.Title,
		Content:   input.Content,
		Author:    input.Author,
		CreatedAt: time.Now(),
	}

	h.posts = append(h.posts, post)
	h.nextID++

	respond(w, http.StatusCreated, post)
}

func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		respondErr(w, http.StatusBadRequest, "id inválido")
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for _, post := range h.posts {
		if post.ID == id {
			respond(w, http.StatusOK, post)
			return
		}
	}

	respondErr(w, http.StatusNotFound, "post não encontrado")
}