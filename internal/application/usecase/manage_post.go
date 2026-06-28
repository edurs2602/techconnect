package usecase

import (
	"context"
	"techconnect/internal/domain/post"
	"time"
)

// --- DTOs ---

type CreatePostInput struct {
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostOutput struct {
	ID        string          `json:"id"`
	UserID    string          `json:"user_id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	CreatedAt time.Time       `json:"created_at"`
	Comments  []CommentOutput `json:"comments,omitempty"`
}

type CreateCommentInput struct {
	PostID  string `json:"post_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

type CommentOutput struct {
	ID        string    `json:"id"`
	PostID    string    `json:"post_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// --- Create Post ---

type CreatePostUseCase struct {
	postService *post.Service
}

func NewCreatePostUseCase(svc *post.Service) *CreatePostUseCase {
	return &CreatePostUseCase{postService: svc}
}

func (uc *CreatePostUseCase) Execute(ctx context.Context, in CreatePostInput) (*PostOutput, error) {
	p, err := uc.postService.CreatePost(ctx, in.UserID, in.Title, in.Content)
	if err != nil {
		return nil, err
	}
	return &PostOutput{
		ID:        p.ID,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
	}, nil
}

// --- Get Post (with nested comments) ---

type GetPostUseCase struct {
	postService *post.Service
}

func NewGetPostUseCase(svc *post.Service) *GetPostUseCase {
	return &GetPostUseCase{postService: svc}
}

func (uc *GetPostUseCase) Execute(ctx context.Context, id string) (*PostOutput, error) {
	p, err := uc.postService.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}

	comments := make([]CommentOutput, len(p.Comments))
	for i, c := range p.Comments {
		comments[i] = CommentOutput{
			ID:        c.ID,
			PostID:    c.PostID,
			UserID:    c.UserID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
		}
	}

	return &PostOutput{
		ID:        p.ID,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		Comments:  comments,
	}, nil
}

// --- List Posts ---

type ListPostsUseCase struct {
	postService *post.Service
}

func NewListPostsUseCase(svc *post.Service) *ListPostsUseCase {
	return &ListPostsUseCase{postService: svc}
}

func (uc *ListPostsUseCase) Execute(ctx context.Context) ([]PostOutput, error) {
	posts, err := uc.postService.ListPosts(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]PostOutput, len(posts))
	for i, p := range posts {
		out[i] = PostOutput{
			ID:        p.ID,
			UserID:    p.UserID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
		}
	}
	return out, nil
}

// --- Delete Post ---

type DeletePostUseCase struct {
	postService *post.Service
}

func NewDeletePostUseCase(svc *post.Service) *DeletePostUseCase {
	return &DeletePostUseCase{postService: svc}
}

func (uc *DeletePostUseCase) Execute(ctx context.Context, id string) error {
	return uc.postService.DeletePost(ctx, id)
}

// --- Add Comment ---

type AddCommentUseCase struct {
	postService *post.Service
}

func NewAddCommentUseCase(svc *post.Service) *AddCommentUseCase {
	return &AddCommentUseCase{postService: svc}
}

func (uc *AddCommentUseCase) Execute(ctx context.Context, in CreateCommentInput) (*CommentOutput, error) {
	c, err := uc.postService.AddComment(ctx, in.PostID, in.UserID, in.Content)
	if err != nil {
		return nil, err
	}
	return &CommentOutput{
		ID:        c.ID,
		PostID:    c.PostID,
		UserID:    c.UserID,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}, nil
}

// --- Delete Comment ---

type DeleteCommentUseCase struct {
	postService *post.Service
}

func NewDeleteCommentUseCase(svc *post.Service) *DeleteCommentUseCase {
	return &DeleteCommentUseCase{postService: svc}
}

func (uc *DeleteCommentUseCase) Execute(ctx context.Context, id string) error {
	return uc.postService.DeleteComment(ctx, id)
}
