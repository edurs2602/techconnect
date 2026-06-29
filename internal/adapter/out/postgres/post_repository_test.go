package postgres_test

import (
	"context"
	"database/sql"
	"testing"

	"techconnect/internal/adapter/out/postgres"
	"techconnect/internal/domain/post"
	"techconnect/internal/domain/user"
	"techconnect/internal/testhelper"

	"github.com/stretchr/testify/assert"
)

func createTestUser(t *testing.T, db *sql.DB) string {
	t.Helper()

	repo := postgres.NewUserRepository(db)

	u := &user.User{
		Username: "user_test",
		Email:    "test@email.com",
		Password: "hash",
	}

	err := repo.Create(context.Background(), u)
	if err != nil {
		t.Fatalf("erro ao criar user: %v", err)
	}

	if u.ID == "" {
		t.Fatal("ID do user vazio")
	}

	return u.ID
}

func TestPostRepository_Create(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	userID := createTestUser(t, db)
	repo := postgres.NewPostRepository(db)

	p := &post.Post{UserID: userID, Title: "titulo", Content: "conteudo"}
	err := repo.Create(context.Background(), p)

	assert.NoError(t, err)
	assert.NotEmpty(t, p.ID)
}

func TestPostRepository_FindByID(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	userID := createTestUser(t, db)
	repo := postgres.NewPostRepository(db)

	p := &post.Post{UserID: userID, Title: "titulo", Content: "conteudo"}
	repo.Create(context.Background(), p)

	found, err := repo.FindByID(context.Background(), p.ID)

	assert.NoError(t, err)
	assert.Equal(t, "titulo", found.Title)
}

func TestPostRepository_List(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	userID := createTestUser(t, db)
	repo := postgres.NewPostRepository(db)

	repo.Create(context.Background(), &post.Post{UserID: userID, Title: "t1", Content: "c1"})
	repo.Create(context.Background(), &post.Post{UserID: userID, Title: "t2", Content: "c2"})

	posts, err := repo.List(context.Background())

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(posts), 2)
}

func TestCommentRepository_CreateAndFind(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	userID := createTestUser(t, db)
	postRepo := postgres.NewPostRepository(db)
	p := &post.Post{UserID: userID, Title: "titulo", Content: "conteudo"}
	postRepo.Create(context.Background(), p)

	commentRepo := postgres.NewCommentRepository(db)
	c := &post.Comment{PostID: p.ID, UserID: userID, Content: "comentario"}
	err := commentRepo.Create(context.Background(), c)

	assert.NoError(t, err)
	assert.NotEmpty(t, c.ID)

	comments, err := commentRepo.FindByPostID(context.Background(), p.ID)

	assert.NoError(t, err)
	assert.Len(t, comments, 1)
	assert.Equal(t, "comentario", comments[0].Content)
}
