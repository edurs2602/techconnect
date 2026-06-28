package post

import "context"

type Service struct {
	postRepo    PostRepository
	commentRepo CommentRepository
}

func NewService(postRepo PostRepository, commentRepo CommentRepository) *Service {
	return &Service{postRepo: postRepo, commentRepo: commentRepo}
}

func (s *Service) CreatePost(ctx context.Context, userID, title, content string) (*Post, error) {
	if userID == "" {
		return nil, ErrorEmptyUserID
	}
	if title == "" {
		return nil, ErrorEmptyTitle
	}
	if content == "" {
		return nil, ErrorEmptyContent
	}

	p := &Post{UserID: userID, Title: title, Content: content}
	if err := s.postRepo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Service) GetPostByID(ctx context.Context, id string) (*Post, error) {
	p, err := s.postRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrorPostNotFound
	}

	comments, err := s.commentRepo.FindByPostID(ctx, id)
	if err != nil {
		return nil, err
	}
	p.Comments = comments

	return p, nil
}

func (s *Service) ListPosts(ctx context.Context) ([]Post, error) {
	return s.postRepo.List(ctx)
}

func (s *Service) DeletePost(ctx context.Context, id string) error {
	return s.postRepo.Delete(ctx, id)
}

func (s *Service) AddComment(ctx context.Context, postID, userID, content string) (*Comment, error) {
	if content == "" {
		return nil, ErrorEmptyContent
	}
	if userID == "" {
		return nil, ErrorEmptyUserID
	}

	if _, err := s.postRepo.FindByID(ctx, postID); err != nil {
		return nil, ErrorPostNotFound
	}

	c := &Comment{PostID: postID, UserID: userID, Content: content}
	if err := s.commentRepo.Create(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.commentRepo.Delete(ctx, id)
}
