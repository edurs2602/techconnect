package postgres

import (
	"context"
	"database/sql"
	db "techconnect/db/sqlc"
	"techconnect/internal/domain/post"

	"github.com/google/uuid"
)

type CommentRepository struct {
	queries *db.Queries
}

func NewCommentRepository(database *sql.DB) *CommentRepository {
	return &CommentRepository{
		queries: db.New(database),
	}
}

func (r *CommentRepository) Create(ctx context.Context, c *post.Comment) error {
	postID, err := uuid.Parse(c.PostID)
	if err != nil {
		return err
	}
	userID, err := uuid.Parse(c.UserID)
	if err != nil {
		return err
	}

	result, err := r.queries.CreateComment(ctx, db.CreateCommentParams{
		PostID:  postID,
		UserID:  userID,
		Content: c.Content,
	})
	if err != nil {
		return err
	}

	c.ID = result.ID.String()
	c.CreatedAt = result.CreatedAt
	return nil
}

func (r *CommentRepository) FindByPostID(ctx context.Context, postID string) ([]post.Comment, error) {
	uid, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}

	results, err := r.queries.GetCommentsByPostID(ctx, uid)
	if err != nil {
		return nil, err
	}

	comments := make([]post.Comment, len(results))
	for i, result := range results {
		comments[i] = post.Comment{
			ID:        result.ID.String(),
			PostID:    result.PostID.String(),
			UserID:    result.UserID.String(),
			Content:   result.Content,
			CreatedAt: result.CreatedAt,
		}
	}
	return comments, nil
}

func (r *CommentRepository) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.queries.DeleteComment(ctx, uid)
}
