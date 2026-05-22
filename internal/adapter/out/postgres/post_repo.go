package postgres

import (
	"context"
	"database/sql"
	db "techconnect/db/sqlc"
	"techconnect/internal/domain/post"

	"github.com/google/uuid"
)

type PostRepository struct {
	queries *db.Queries
}

func NewPostRepository(database *sql.DB) *PostRepository {
	return &PostRepository{
		queries: db.New(database),
	}
}

func (r *PostRepository) Create(ctx context.Context, p *post.Post) error {
	uid, err := uuid.Parse(p.UserID)
	if err != nil {
		return err
	}

	result, err := r.queries.CreatePost(ctx, db.CreatePostParams{
		UserID:  uid,
		Title:   p.Title,
		Content: p.Content,
	})
	if err != nil {
		return err
	}

	p.ID = result.ID.String()
	p.CreatedAt = result.CreatedAt
	return nil
}

func (r *PostRepository) FindByID(ctx context.Context, id string) (*post.Post, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	result, err := r.queries.GetPostByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &post.Post{
		ID:        result.ID.String(),
		UserID:    result.UserID.String(),
		Title:     result.Title,
		Content:   result.Content,
		CreatedAt: result.CreatedAt,
	}, nil
}

func (r *PostRepository) List(ctx context.Context) ([]post.Post, error) {
	results, err := r.queries.ListPosts(ctx)
	if err != nil {
		return nil, err
	}

	posts := make([]post.Post, len(results))
	for i, result := range results {
		posts[i] = post.Post{
			ID:        result.ID.String(),
			UserID:    result.UserID.String(),
			Title:     result.Title,
			Content:   result.Content,
			CreatedAt: result.CreatedAt,
		}
	}
	return posts, nil
}

func (r *PostRepository) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.queries.DeletePost(ctx, uid)
}
