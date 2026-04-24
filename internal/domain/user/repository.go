package user

import "context"

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}
