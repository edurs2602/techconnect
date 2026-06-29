package usecase

import (
	"context"
	"errors"
	"techconnect/internal/adapter/out/postgres"
	"time"
)

type RefreshInput struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshUseCase struct {
	tokenRepo *postgres.RefreshTokenRepository
	issuer    TokenIssuer
	validator interface{ Validate(string) (string, error) }
}

func NewRefreshUseCase(repo *postgres.RefreshTokenRepository, issuer TokenIssuer, validator interface{ Validate(string) (string, error) }) *RefreshUseCase {
	return &RefreshUseCase{tokenRepo: repo, issuer: issuer, validator: validator}
}

func (uc *RefreshUseCase) Execute(ctx context.Context, in RefreshInput) (*LoginOutput, error) {
	userID, err := uc.validator.Validate(in.RefreshToken)
	if err != nil {
		return nil, errors.New("token inválido")
	}

	stored, err := uc.tokenRepo.Get(ctx, in.RefreshToken)
	if err != nil || stored.Used || stored.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expirado ou já utilizado")
	}

	// rotação
	uc.tokenRepo.MarkUsed(ctx, in.RefreshToken)

	access, _ := uc.issuer.IssueAccess(userID)
	refresh, _ := uc.issuer.IssueRefresh(userID)

	uc.tokenRepo.Save(ctx, userID, refresh, time.Now().Add(7*24*time.Hour))

	return &LoginOutput{AccessToken: access, RefreshToken: refresh}, nil
}
