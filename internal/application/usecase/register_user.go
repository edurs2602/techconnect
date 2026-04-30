package usecase

import (
	"context"

	"techconnect/internal/domain/user"
)

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterUseCase struct {
	userService *user.Service
}

func NewRegisterUseCase(svc *user.Service) *RegisterUseCase {
	return &RegisterUseCase{userService: svc}
}

func (uc *RegisterUseCase) Execute(ctx context.Context, in RegisterInput) (*RegisterOutput, error) {
	u, err := uc.userService.Register(ctx, in.Username, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	return &RegisterOutput{ID: u.ID, Username: u.Username, Email: u.Email}, nil
}
