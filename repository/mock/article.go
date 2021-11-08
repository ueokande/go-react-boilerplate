package mock

import (
	"context"
	"time"

	"github.com/ueokande/go-react-builderplate/repository"
)

func (repo Repository) GetAllArticles(ctx context.Context) ([]repository.Article, error) {
	mu.RLock()
	defer mu.RUnlock()

	return mockArticles, nil
}

func (repo Repository) GetAllArticle(ctx context.Context, id int64) (repository.Article, error) {
	for _, a := range mockArticles {
		if a.ID == id {
			return a, nil
		}
	}
	return repository.Article{}, repository.ErrArticleNotFound
}

func (repo Repository) SaveArticle(ctx context.Context, title, content, author string, createdAt time.Time) (repository.Article, error) {
	mu.Lock()
	defer mu.Unlock()

	var article repository.Article
	article.ID = articleIDCounter
	article.Title = title
	article.Content = content
	article.Author = author
	article.CreatedAt = createdAt

	mockArticles = append(mockArticles, article)
	articleIDCounter++

	return article, nil
}
