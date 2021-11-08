package mock

import (
	"context"
	"sort"
	"time"

	"github.com/ueokande/go-react-builderplate/repository"
)

func (repo Repository) GetComments(ctx context.Context, articleID int64) ([]repository.Comment, error) {
	mu.RLock()
	defer mu.RUnlock()

	_, err := repo.GetAllArticle(ctx, articleID)
	if err != nil {
		return nil, err
	}

	var comments []repository.Comment
	for _, c := range mockComments {
		if c.ArticleID == articleID {
			comments = append(comments, c)
		}
	}
	sort.Slice(comments, func(i, j int) bool { return comments[i].CreatedAt.After(comments[j].CreatedAt) })

	return comments, nil
}

func (repo Repository) SaveComment(ctx context.Context, articleID int64, author string, content string, createdAt time.Time) (repository.Comment, error) {
	mu.Lock()
	defer mu.Unlock()

	_, err := repo.GetAllArticle(ctx, articleID)
	if err != nil {
		return repository.Comment{}, err
	}

	var comment repository.Comment
	comment.ID = commentIDCounter
	comment.ArticleID = articleID
	comment.Content = content
	comment.Author = author
	comment.CreatedAt = createdAt

	mockComments = append(mockComments, comment)
	commentIDCounter++

	return comment, nil
}
