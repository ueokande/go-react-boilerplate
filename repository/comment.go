package repository

import (
	"context"
	"time"
)

type Comment struct {
	ID        int64
	ArticleID int64
	Author    string
	Content   string
	CreatedAt time.Time
}

type CommentRepository interface {
	GetComments(ctx context.Context, articleID int64) ([]Comment, error)
	SaveComment(ctx context.Context, articleID int64, author string, content string, createdAt time.Time) (Comment, error)
}
