package postgres

import (
	"context"
	"database/sql"
	db "techconnect/db/sqlc"
	"time"

	"github.com/google/uuid"
)

type RefreshTokenRepository struct {
	queries *db.Queries
}

func NewRefreshTokenRepository(database *sql.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{queries: db.New(database)}
}

type StoredToken struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
	Used      bool
}

func (r *RefreshTokenRepository) Save(ctx context.Context, userID, token string, expiresAt time.Time) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	_, err = r.queries.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		UserID:    uid,
		Token:     token,
		ExpiresAt: expiresAt,
	})
	return err
}

func (r *RefreshTokenRepository) Get(ctx context.Context, token string) (*StoredToken, error) {
	result, err := r.queries.GetRefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return &StoredToken{
		UserID:    result.UserID.String(),
		Token:     result.Token,
		ExpiresAt: result.ExpiresAt,
		Used:      result.Used,
	}, nil
}

func (r *RefreshTokenRepository) MarkUsed(ctx context.Context, token string) error {
	return r.queries.MarkTokenUsed(ctx, token)
}

func (r *RefreshTokenRepository) DeleteByUser(ctx context.Context, userID string) error {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	return r.queries.DeleteUserTokens(ctx, uid)
}
