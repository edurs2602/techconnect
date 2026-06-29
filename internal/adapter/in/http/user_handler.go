package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"techconnect/internal/application/usecase"
	"techconnect/internal/domain/user"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	register   *usecase.RegisterUseCase
	login      *usecase.LoginUseCase
	getUser    *usecase.GetUserUseCase
	updateUser *usecase.UpdateUserUseCase
	deleteUser *usecase.DeleteUserUseCase
	refresh    *usecase.RefreshUseCase
}

func NewUserHandler(r *usecase.RegisterUseCase, l *usecase.LoginUseCase, g *usecase.GetUserUseCase, u *usecase.UpdateUserUseCase, d *usecase.DeleteUserUseCase, rf *usecase.RefreshUseCase) *UserHandler {

	return &UserHandler{
		register:   r,
		login:      l,
		getUser:    g,
		updateUser: u,
		deleteUser: d,
		refresh:    rf,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	out, err := h.getUser.Execute(r.Context(), username)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrorUserNotFound):
			respondErr(w, http.StatusNotFound, err.Error())
		default:
			respondErr(w, http.StatusInternalServerError, "erro interno")
		}
		return
	}

	respond(w, http.StatusOK, out)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	var in usecase.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondErr(w, http.StatusBadRequest, "payload inválido")
		return
	}

	out, err := h.updateUser.Execute(r.Context(), userID, in)
	if err != nil {
		respondErr(w, http.StatusInternalServerError, "erro interno")
		return
	}

	respond(w, http.StatusOK, out)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var in usecase.RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondErr(w, http.StatusBadRequest, "payload inválido")
		return
	}

	out, err := h.register.Execute(r.Context(), in)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrorEmptyUsername),
			errors.Is(err, user.ErrorEmptyEmail),
			errors.Is(err, user.ErrorEmptyPassword):
			respondErr(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, user.ErrorEmailTaken),
			errors.Is(err, user.ErrorUsernameTaken):
			respondErr(w, http.StatusConflict, err.Error())

		default:
			respondErr(w, http.StatusInternalServerError, "erro interno")
		}
		return
	}

	respond(w, http.StatusCreated, out)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	if err := h.deleteUser.Execute(r.Context(), userID); err != nil {
		respondErr(w, http.StatusInternalServerError, "erro interno")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var in usecase.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondErr(w, http.StatusBadRequest, "payload inválido")
		return
	}

	out, err := h.login.Execute(r.Context(), in)
	if err != nil {
		respondErr(w, http.StatusUnauthorized, "credenciais inválidas")
		return
	}

	respond(w, http.StatusOK, out)
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var in usecase.RefreshInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondErr(w, http.StatusBadRequest, "payload inválido")
		return
	}

	out, err := h.refresh.Execute(r.Context(), in)
	if err != nil {
		respondErr(w, http.StatusUnauthorized, err.Error())
		return
	}

	respond(w, http.StatusOK, out)
}
