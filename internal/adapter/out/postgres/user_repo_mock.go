package postgres

//import (
//	"context"
//	"sync"
//	"techconnect/internal/domain/user"
//
//	"github.com/google/uuid"
//)

//type UserRepository struct {
//	mu    sync.RWMutex
//	users map[string]*user.User // chave = ID
//}

//func NewUserRepository() *UserRepository {
//	return &UserRepository{
//		users: make(map[string]*user.User),
//	}
//}

//func (r *UserRepository) Create(ctx context.Context, u *user.User) error {
//	r.mu.Lock()
//	defer r.mu.Unlock()

//	u.ID = uuid.NewString()
//	r.users[u.ID] = u
//	return nil
//}

//func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
//	r.mu.RLock()
//	defer r.mu.RUnlock()

//	for _, u := range r.users {
//		if u.Email == email {
//			return true, nil
//		}
//	}
//	return false, nil
//}

//func (r *UserRepository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
//	r.mu.RLock()
//	defer r.mu.RUnlock()

//	for _, u := range r.users {
//		if u.Username == username {
//			return true, nil
//		}
//	}
//	return false, nil
//}
