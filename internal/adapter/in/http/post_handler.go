package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"techconnect/internal/application/usecase"
	"techconnect/internal/domain/post"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	createPost    *usecase.CreatePostUseCase
	getPost       *usecase.GetPostUseCase
	listPosts     *usecase.ListPostsUseCase
	deletePost    *usecase.DeletePostUseCase
	addComment    *usecase.AddCommentUseCase
	deleteComment *usecase.DeleteCommentUseCase
}

func NewPostHandler(
	cp *usecase.CreatePostUseCase,
	gp *usecase.GetPostUseCase,
	lp *usecase.ListPostsUseCase,
	dp *usecase.DeletePostUseCase,
	ac *usecase.AddCommentUseCase,
	dc *usecase.DeleteCommentUseCase,
) *PostHandler {
	return &PostHandler{
		createPost:    cp,
		getPost:       gp,
		listPosts:     lp,
		deletePost:    dp,
		addComment:    ac,
		deleteComment: dc,
	}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in usecase.CreatePostInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondErr(w, http.StatusBadRequest, "payload inválido")
		return
	}

	out, err := h.createPost.Execute(r.Context(), in)
	if err != nil {
		switch {
		case errors.Is(err, post.ErrorEmptyTitle),
			errors.Is(err, post.ErrorEmptyContent),
			errors.Is(err, post.ErrorEmptyUserID):
			respondErr(w, http.StatusBadRequest, err.Error())
		default:
			respondErr(w, http.StatusInternalServerError, "erro interno")
		}
		return
	}

	respond(w, http.StatusCreated, out)
}

func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	out, err := h.getPost.Execute(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, post.ErrorPostNotFound):
			respondErr(w, http.StatusNotFound, err.Error())
		default:
			respondErr(w, http.StatusInternalServerError, "erro interno")
		}
		return
	}

	respond(w, http.StatusOK, out)
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	out, err := h.listPosts.Execute(r.Context())
	if err != nil {
		respondErr(w, http.StatusInternalServerError, "erro interno")
		return
	}

	respond(w, http.StatusOK, out)
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.deletePost.Execute(r.Context(), id); err != nil {
		respondErr(w, http.StatusInternalServerError, "erro interno")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PostHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")

	var in usecase.CreateCommentInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondErr(w, http.StatusBadRequest, "payload inválido")
		return
	}
	in.PostID = postID

	out, err := h.addComment.Execute(r.Context(), in)
	if err != nil {
		switch {
		case errors.Is(err, post.ErrorPostNotFound):
			respondErr(w, http.StatusNotFound, err.Error())
		case errors.Is(err, post.ErrorEmptyContent),
			errors.Is(err, post.ErrorEmptyUserID):
			respondErr(w, http.StatusBadRequest, err.Error())
		default:
			respondErr(w, http.StatusInternalServerError, "erro interno")
		}
		return
	}

	respond(w, http.StatusCreated, out)
}

func (h *PostHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID := chi.URLParam(r, "commentId")

	if err := h.deleteComment.Execute(r.Context(), commentID); err != nil {
		respondErr(w, http.StatusInternalServerError, "erro interno")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
