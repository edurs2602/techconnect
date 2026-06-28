package usecase

import (
	"context"
	"errors"
	"techconnect/internal/domain/user"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Authenticator interface {
	Compare(hashed, plain string) error
}

type TokenIssuer interface {
	IssueAccess(userID string) (string, error)
	IssueRefresh(userID string) (string, error)
}

type LoginUseCase struct {
	userService *user.Service
	auth        Authenticator
	issuer      TokenIssuer
}

func NewLoginUseCase(svc *user.Service, auth Authenticator, issuer TokenIssuer) *LoginUseCase {
	return &LoginUseCase{userService: svc, auth: auth, issuer: issuer}
}

func (uc *LoginUseCase) Execute(ctx context.Context, in LoginInput) (*LoginOutput, error) {
	u, err := uc.userService.GetByEmail(ctx, in.Email)
	if err != nil {
		return nil, user.ErrorUserNotFound
	}

	if err := uc.auth.Compare(u.Password, in.Password); err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	access, err := uc.issuer.IssueAccess(u.ID)
	if err != nil {
		return nil, err
	}

	refresh, err := uc.issuer.IssueRefresh(u.ID)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{AccessToken: access, RefreshToken: refresh}, nil
}
