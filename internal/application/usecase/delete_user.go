package usecase

import (
	"context"
	"techconnect/internal/domain/user"
)

type DeleteUserUseCase struct {
	userService *user.Service
}

func NewDeleteUserUseCase(svc *user.Service) *DeleteUserUseCase {
	return &DeleteUserUseCase{userService: svc}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, id string) error {
	return uc.userService.Delete(ctx, id)
}
