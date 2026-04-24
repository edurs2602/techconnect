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
