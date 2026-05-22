package post

import (
	"errors"
	"time"
)

type Post struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
	Comments  []Comment
}

type Comment struct {
	ID        string
	PostID    string
	UserID    string
	Content   string
	CreatedAt time.Time
}

var (
	ErrorPostNotFound    = errors.New("post não encontrado")
	ErrorEmptyTitle      = errors.New("título obrigatório")
	ErrorEmptyContent    = errors.New("conteúdo obrigatório")
	ErrorEmptyUserID     = errors.New("usuário obrigatório")
	ErrorCommentNotFound = errors.New("comentário não encontrado")
)
