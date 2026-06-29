package postgres_test

import (
	"context"
	"testing"

	"techconnect/internal/adapter/out/postgres"
	"techconnect/internal/domain/user"
	"techconnect/internal/testhelper"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	repo := postgres.NewUserRepository(db)

	u := &user.User{
		Username: "joao",
		Email:    "joao@email.com",
		Password: "hashed_senha",
	}

	err := repo.Create(context.Background(), u)

	assert.NoError(t, err)
	assert.NotEmpty(t, u.ID)
}

func TestUserRepository_ExistsByEmail_True(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	repo := postgres.NewUserRepository(db)
	repo.Create(context.Background(), &user.User{
		Username: "joao",
		Email:    "joao@email.com",
		Password: "hash",
	})

	exists, err := repo.ExistsByEmail(context.Background(), "joao@email.com")

	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestUserRepository_ExistsByEmail_False(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	repo := postgres.NewUserRepository(db)

	exists, err := repo.ExistsByEmail(context.Background(), "naoexiste@email.com")

	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	repo := postgres.NewUserRepository(db)
	repo.Create(context.Background(), &user.User{
		Username: "joao",
		Email:    "joao@email.com",
		Password: "hash",
	})

	u, err := repo.FindByUsername(context.Background(), "joao")

	assert.NoError(t, err)
	assert.Equal(t, "joao", u.Username)
}

func TestUserRepository_FindByUsername_NaoEncontrado(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	repo := postgres.NewUserRepository(db)

	u, err := repo.FindByUsername(context.Background(), "naoexiste")

	assert.Nil(t, u)
	assert.Error(t, err)
}

func TestUserRepository_Update(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	repo := postgres.NewUserRepository(db)
	u := &user.User{Username: "joao", Email: "joao@email.com", Password: "hash"}
	repo.Create(context.Background(), u)

	updated, err := repo.UpdateUser(context.Background(), &user.User{
		ID:       u.ID,
		Username: "joao_novo",
		Bio:      "minha bio",
	})

	assert.NoError(t, err)
	assert.Equal(t, "joao_novo", updated.Username)
	assert.Equal(t, "minha bio", updated.Bio)
}

func TestUserRepository_Delete(t *testing.T) {
	db := testhelper.NewTestDB(t)
	defer testhelper.CleanDB(t, db)

	repo := postgres.NewUserRepository(db)
	u := &user.User{Username: "joao", Email: "joao@email.com", Password: "hash"}
	repo.Create(context.Background(), u)

	err := repo.DeleteUser(context.Background(), u.ID)

	assert.NoError(t, err)

	exists, _ := repo.ExistsByEmail(context.Background(), "joao@email.com")
	assert.False(t, exists)
}
