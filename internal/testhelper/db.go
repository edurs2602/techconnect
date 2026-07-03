package testhelper

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func NewTestDB(t *testing.T) *sql.DB {
	t.Helper()

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgresql://social:social@localhost:5433/social?sslmode=disable"
	}

	db, err := sql.Open("postgres", url)
	if err != nil {
		t.Fatalf("erro ao conectar no banco: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("banco inacessível: %v", err)
	}

	return db
}

func CleanDB(t *testing.T, db *sql.DB) {
	t.Helper()

	db.Exec("DELETE FROM comments")
	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM refresh_tokens")
	db.Exec("DELETE FROM users")
}
