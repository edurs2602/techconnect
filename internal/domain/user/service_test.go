package user

import (
	"context"
	"errors"
	"testing"
)

// --- mocks ---

type mockRepo struct {
	existsByEmail    bool
	existsByUsername bool
	createErr        error
}

func (m *mockRepo) Create(_ context.Context, u *User) error {
	if m.createErr != nil {
		return m.createErr
	}
	u.ID = "fake-uuid"
	return nil
}

func (m *mockRepo) ExistsByEmail(_ context.Context, _ string) (bool, error) {
	return m.existsByEmail, nil
}

func (m *mockRepo) ExistsByUsername(_ context.Context, _ string) (bool, error) {
	return m.existsByUsername, nil
}

type mockHasher struct {
	err error
}

func (m *mockHasher) Hash(plain string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return "hashed-" + plain, nil
}

// --- tests ---

func TestRegister_Success(t *testing.T) {
	svc := NewService(&mockRepo{}, &mockHasher{})

	u, err := svc.Register(context.Background(), "john", "john@example.com", "secret")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.ID != "fake-uuid" {
		t.Errorf("expected ID fake-uuid, got %s", u.ID)
	}
	if u.Username != "john" {
		t.Errorf("expected username john, got %s", u.Username)
	}
	if u.Email != "john@example.com" {
		t.Errorf("expected email john@example.com, got %s", u.Email)
	}
	if u.Password != "hashed-secret" {
		t.Errorf("expected hashed password, got %s", u.Password)
	}
}

func TestRegister_EmptyUsername(t *testing.T) {
	svc := NewService(&mockRepo{}, &mockHasher{})

	_, err := svc.Register(context.Background(), "", "john@example.com", "secret")
	if !errors.Is(err, ErrorEmptyUsername) {
		t.Errorf("expected ErrorEmptyUsername, got %v", err)
	}
}

func TestRegister_EmptyEmail(t *testing.T) {
	svc := NewService(&mockRepo{}, &mockHasher{})

	_, err := svc.Register(context.Background(), "john", "", "secret")
	if !errors.Is(err, ErrorEmptyEmail) {
		t.Errorf("expected ErrorEmptyEmail, got %v", err)
	}
}

func TestRegister_EmptyPassword(t *testing.T) {
	svc := NewService(&mockRepo{}, &mockHasher{})

	_, err := svc.Register(context.Background(), "john", "john@example.com", "")
	if !errors.Is(err, ErrorEmptyPassword) {
		t.Errorf("expected ErrorEmptyPassword, got %v", err)
	}
}

func TestRegister_EmailTaken(t *testing.T) {
	svc := NewService(&mockRepo{existsByEmail: true}, &mockHasher{})

	_, err := svc.Register(context.Background(), "john", "john@example.com", "secret")
	if !errors.Is(err, ErrorEmailTaken) {
		t.Errorf("expected ErrorEmailTaken, got %v", err)
	}
}

func TestRegister_UsernameTaken(t *testing.T) {
	svc := NewService(&mockRepo{existsByUsername: true}, &mockHasher{})

	_, err := svc.Register(context.Background(), "john", "john@example.com", "secret")
	if !errors.Is(err, ErrorUsernameTaken) {
		t.Errorf("expected ErrorUsernameTaken, got %v", err)
	}
}

func TestRegister_HasherError(t *testing.T) {
	hashErr := errors.New("hash failure")
	svc := NewService(&mockRepo{}, &mockHasher{err: hashErr})

	_, err := svc.Register(context.Background(), "john", "john@example.com", "secret")
	if !errors.Is(err, hashErr) {
		t.Errorf("expected hash error, got %v", err)
	}
}

func TestRegister_RepoCreateError(t *testing.T) {
	repoErr := errors.New("db error")
	svc := NewService(&mockRepo{createErr: repoErr}, &mockHasher{})

	_, err := svc.Register(context.Background(), "john", "john@example.com", "secret")
	if !errors.Is(err, repoErr) {
		t.Errorf("expected repo error, got %v", err)
	}
}
