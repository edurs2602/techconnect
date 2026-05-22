package usecase

import (
	"context"
	"techconnect/internal/domain/user"
)

type GetUserOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type GetUserUseCase struct {
	userService *user.Service
}

func NewGetUserUseCase(svc *user.Service) *GetUserUseCase {
	return &GetUserUseCase{userService: svc}
}

func (uc *GetUserUseCase) Execute(ctx context.Context, username string) (*GetUserOutput, error) {
	u, err := uc.userService.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &GetUserOutput{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Bio:      u.Bio,
	}, nil
}
