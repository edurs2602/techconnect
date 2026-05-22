package user

import "context"

type Hasher interface {
	Hash(plain string) (string, error)
}

type Service struct {
	repo   UserRepository
	hasher Hasher
}

func NewService(repo UserRepository, hasher Hasher) *Service {
	return &Service{repo: repo, hasher: hasher}
}

func (s *Service) Register(ctx context.Context, username, email, password string) (*User, error) {

	if username == "" {
		return nil, ErrorEmptyUsername
	}
	if email == "" {
		return nil, ErrorEmptyEmail
	}
	if password == "" {
		return nil, ErrorEmptyPassword
	}

	if ok, _ := s.repo.ExistsByEmail(ctx, email); ok {
		return nil, ErrorEmailTaken
	}
	if ok, _ := s.repo.ExistsByUsername(ctx, username); ok {
		return nil, ErrorUsernameTaken
	}

	hash, err := s.hasher.Hash(password)
	if err != nil {
		return nil, err
	}

	u := &User{Username: username, Email: email, Password: hash}
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*User, error) {
	u, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, ErrorUserNotFound
	}
	return u, nil
}

func (s *Service) Update(ctx context.Context, id, username, bio string) (*User, error) {
	u := &User{ID: id, Username: username, Bio: bio}
	updated, err := s.repo.UpdateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
