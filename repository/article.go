package repository

import (
	"context"
	"time"
)

type Article struct {
	ID        int64
	Title     string
	Content   string
	CreatedAt time.Time
	Author    string
}

type ArticleRepository interface {
	GetAllArticle(ctx context.Context, id int64) (Article, error)
	GetAllArticles(ctx context.Context) ([]Article, error)
	SaveArticle(ctx context.Context, title, content, author string, createdAt time.Time) (Article, error)
}
