package post

import "context"

type PostRepository interface {
	Create(ctx context.Context, p *Post) error
	FindByID(ctx context.Context, id string) (*Post, error)
	List(ctx context.Context) ([]Post, error)
	Delete(ctx context.Context, id string) error
}

type CommentRepository interface {
	Create(ctx context.Context, c *Comment) error
	FindByPostID(ctx context.Context, postID string) ([]Comment, error)
	Delete(ctx context.Context, id string) error
}
