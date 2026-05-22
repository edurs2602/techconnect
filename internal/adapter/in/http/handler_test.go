package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"techconnect/internal/application/usecase"
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

func newTestHandler() *UserHandler {
	svc := user.NewService(&mockRepo{}, &mockHasher{})
	uc := usecase.NewRegisterUseCase(svc)
	return NewUserHandler(uc)
}

// --- Register tests ---

func TestRegister_Created(t *testing.T) {
	h := newTestHandler()
	body := `{"username":"john","email":"john@example.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.Register(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}

	var out usecase.RegisterOutput
	json.NewDecoder(rec.Body).Decode(&out)
	if out.ID != "test-id" {
		t.Errorf("expected ID test-id, got %s", out.ID)
	}
}

func TestRegister_InvalidJSON(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString("{bad"))
	rec := httptest.NewRecorder()

	h.Register(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

func TestRegister_EmptyUsername(t *testing.T) {
	h := newTestHandler()
	body := `{"username":"","email":"john@example.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.Register(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

func TestRegister_EmptyEmail(t *testing.T) {
	h := newTestHandler()
	body := `{"username":"john","email":"","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.Register(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

func TestRegister_EmptyPassword(t *testing.T) {
	h := newTestHandler()
	body := `{"username":"john","email":"john@example.com","password":""}`
	req := httptest.NewRequest(http.MethodPost, "/auth/register", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.Register(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

// --- Login tests ---

func TestLogin_Success(t *testing.T) {
	h := newTestHandler()
	body := `{"email":"john@example.com","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.Login(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var resp map[string]string
	json.NewDecoder(rec.Body).Decode(&resp)
	if resp["token"] == "" {
		t.Error("expected token in response")
	}
}

func TestLogin_InvalidJSON(t *testing.T) {
	h := newTestHandler()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString("{bad"))
	rec := httptest.NewRecorder()

	h.Login(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

func TestLogin_EmptyEmail(t *testing.T) {
	h := newTestHandler()
	body := `{"email":"","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.Login(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

func TestLogin_EmptyPassword(t *testing.T) {
	h := newTestHandler()
	body := `{"email":"john@example.com","password":""}`
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	h.Login(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}
