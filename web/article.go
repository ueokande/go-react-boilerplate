package web

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/ueokande/go-react-builderplate/repository"
)

func (h APIHandler) GetArticleSummaries(w http.ResponseWriter, r *http.Request) {
	articles, err := h.ArticleRepository.GetAllArticles(r.Context())
	if err != nil {
		log.Print(err)
		InternalServerError(w)
		return
	}

	type ArticleJSON = struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		Summary   string    `json:"summary"`
		CreatedAt time.Time `json:"created_at"`
		Author    string    `json:"author"`
	}
	var out struct {
		Articles []ArticleJSON `json:"article_summaries"`
	}
	out.Articles = make([]ArticleJSON, len(articles))
	for i, a := range articles {
		out.Articles[i].ID = a.ID
		out.Articles[i].Title = a.Title
		out.Articles[i].Summary = strings.Split(a.Content, "\n\n")[0]
		out.Articles[i].CreatedAt = a.CreatedAt
		out.Articles[i].Author = a.Author
	}
	ResponseJSON(w, out, http.StatusOK)
}

func (h APIHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["article_id"], 10, 64)
	if err != nil {
		ErrorJSON(w, "invalid article id", http.StatusBadRequest)
		return
	}
	article, err := h.ArticleRepository.GetAllArticle(r.Context(), id)
	if errors.Is(err, repository.ErrArticleNotFound) {
		ErrorJSON(w, "article not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Print(err)
		InternalServerError(w)
		return
	}

	var out struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
		Author    string    `json:"author"`
	}
	out.ID = article.ID
	out.Title = article.Title
	out.Content = article.Content
	out.CreatedAt = article.CreatedAt
	out.Author = article.Author
	ResponseJSON(w, out, http.StatusOK)
}

func (h APIHandler) GetArticleComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID, err := strconv.ParseInt(vars["article_id"], 10, 64)
	if err != nil {
		ErrorJSON(w, "invalid article id", http.StatusBadRequest)
		return
	}

	comments, err := h.CommentRepository.GetComments(r.Context(), articleID)
	if errors.Is(err, repository.ErrArticleNotFound) {
		ErrorJSON(w, "article not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Print(err)
		InternalServerError(w)
		return
	}

	type CommentJSON struct {
		ID        int64     `json:"id"`
		ArticleID int64     `json:"article_id"`
		Author    string    `json:"author"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}
	var out struct {
		Comments []CommentJSON `json:"comments"`
	}
	out.Comments = make([]CommentJSON, len(comments))
	for i, c := range comments {
		out.Comments[i].ID = c.ID
		out.Comments[i].ArticleID = c.ID
		out.Comments[i].Author = c.Author
		out.Comments[i].Content = c.Content
		out.Comments[i].CreatedAt = c.CreatedAt
	}
	ResponseJSON(w, out, http.StatusOK)
}

func (h APIHandler) PostArticle(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
		Author    string    `json:"author"`
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		return
	}
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Print(err)
		ErrorJSON(w, "invalid json", http.StatusBadRequest)
		return
	}
	if input.Title == "" || input.Content == "" || input.Author == "" {
		ErrorJSON(w, "invalid content", http.StatusBadRequest)
		return
	}

	article, err := h.ArticleRepository.SaveArticle(r.Context(),
		input.Title,
		input.Content,
		input.Author,
		time.Now(),
	)
	if err != nil {
		log.Print(err)
		InternalServerError(w)
		return
	}

	var out struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
		Author    string    `json:"author"`
	}
	out.ID = article.ID
	out.Title = article.Title
	out.Content = article.Content
	out.CreatedAt = article.CreatedAt
	out.Author = article.Author
	ResponseJSON(w, out, http.StatusCreated)
}

func (h APIHandler) PostArticleComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleID, err := strconv.ParseInt(vars["article_id"], 10, 64)
	if err != nil {
		ErrorJSON(w, "invalid article id", http.StatusBadRequest)
		return
	}
	var input struct {
		Author  string `json:"author"`
		Content string `json:"content"`
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		return
	}
	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Print(err)
		ErrorJSON(w, "invalid json", http.StatusBadRequest)
		return
	}
	if input.Author == "" || input.Content == "" {
		ErrorJSON(w, "invalid content", http.StatusBadRequest)
		return
	}

	comment, err := h.CommentRepository.SaveComment(r.Context(), articleID, input.Author, input.Content, time.Now())
	if errors.Is(err, repository.ErrArticleNotFound) {
		ErrorJSON(w, "article not found", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Print(err)
		InternalServerError(w)
		return
	}

	var out struct {
		ID        int64     `json:"id"`
		ArticleID int64     `json:"article_id"`
		Author    string    `json:"author"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}
	out.ID = comment.ID
	out.ArticleID = comment.ID
	out.Author = comment.Author
	out.Content = comment.Content
	out.CreatedAt = comment.CreatedAt
	ResponseJSON(w, out, http.StatusCreated)
}
