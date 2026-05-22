package user

import "context"

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}
