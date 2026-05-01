package usecase

import (
	"context"
	"testing"

	"techconnect/internal/domain/user"
)

// --- mocks ---

type mockRepo struct{}

func (m *mockRepo) Create(_ context.Context, u *user.User) error {
	u.ID = "test-id"
	return nil
}
func (m *mockRepo) ExistsByEmail(_ context.Context, _ string) (bool, error)    { return false, nil }
func (m *mockRepo) ExistsByUsername(_ context.Context, _ string) (bool, error) { return false, nil }

type mockHasher struct{}

func (m *mockHasher) Hash(plain string) (string, error) { return "hashed", nil }

// --- tests ---

func TestRegisterUseCase_Success(t *testing.T) {
	svc := user.NewService(&mockRepo{}, &mockHasher{})
	uc := NewRegisterUseCase(svc)

	out, err := uc.Execute(context.Background(), RegisterInput{
		Username: "john",
		Email:    "john@example.com",
		Password: "secret",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if out.ID != "test-id" {
		t.Errorf("expected ID test-id, got %s", out.ID)
	}
	if out.Username != "john" {
		t.Errorf("expected username john, got %s", out.Username)
	}
	if out.Email != "john@example.com" {
		t.Errorf("expected email john@example.com, got %s", out.Email)
	}
}

func TestRegisterUseCase_ValidationError(t *testing.T) {
	svc := user.NewService(&mockRepo{}, &mockHasher{})
	uc := NewRegisterUseCase(svc)

	_, err := uc.Execute(context.Background(), RegisterInput{
		Username: "",
		Email:    "john@example.com",
		Password: "secret",
	})
	if err == nil {
		t.Fatal("expected error for empty username")
	}
}
