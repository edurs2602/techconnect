package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"techconnect/internal/application/usecase"
	"techconnect/internal/domain/user"
)

type UserHandler struct {
	register *usecase.RegisterUseCase
}

func NewUserHandler(r *usecase.RegisterUseCase) *UserHandler {
	return &UserHandler{register: r}
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
			respondErr(w, http.StatusConflict, err.Error()) // 409
		default:
			respondErr(w, http.StatusInternalServerError, "erro interno")
		}
		return
	}

	respond(w, http.StatusCreated, out)
}
