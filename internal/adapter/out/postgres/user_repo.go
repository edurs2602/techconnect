package postgres

import (
	"context"
	"database/sql"
	db "techconnect/db/sqlc"
	"techconnect/internal/domain/user"

	"github.com/google/uuid"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		queries: db.New(database),
	}
}

func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
	result, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return err
	}
	u.ID = result.ID.String()
	return nil
}

func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return r.queries.ExistsByEmail(ctx, email)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	result, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:       result.ID.String(),
		Username: result.Username,
		Email:    result.Email,
		Password: result.Password,
		Bio:      result.Bio.String,
	}, nil
}

func (r *UserRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	return r.queries.ExistsByUsername(ctx, username)
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*user.User, error) {
	result, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:       result.ID.String(),
		Username: result.Username,
		Email:    result.Email,
		Password: result.Password,
		Bio:      result.Bio.String,
	}, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, u *user.User) (*user.User, error) {
	uid, err := uuid.Parse(u.ID)
	if err != nil {
		return nil, err
	}

	result, err := r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:       uid,
		Username: u.Username,
		Bio:      sql.NullString{String: u.Bio, Valid: u.Bio != ""},
	})
	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:       result.ID.String(),
		Username: result.Username,
		Email:    result.Email,
		Bio:      result.Bio.String,
	}, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.queries.DeleteUser(ctx, uid)
}
