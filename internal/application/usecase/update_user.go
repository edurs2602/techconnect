package usecase

import (
	"context"
	"techconnect/internal/domain/user"
)

type UpdateUserInput struct {
	Username string `json:"username"`
	Bio      string `json:"bio"`
}

type UpdateUserUseCase struct {
	userService *user.Service
}

func NewUpdateUserUseCase(svc *user.Service) *UpdateUserUseCase {
	return &UpdateUserUseCase{userService: svc}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, id string, in UpdateUserInput) (*GetUserOutput, error) {
	u, err := uc.userService.Update(ctx, id, in.Username, in.Bio)
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
