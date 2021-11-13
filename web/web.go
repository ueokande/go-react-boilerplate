package web

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ueokande/go-react-builderplate/repository"
)

type APIHandler struct {
	ArticleRepository repository.ArticleRepository
	CommentRepository repository.CommentRepository
}

func NewWeb(asset http.Handler, api APIHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/health", api.Health).Methods(http.MethodGet)
	r.HandleFunc("/api/ready", api.Ready).Methods(http.MethodGet)
	r.HandleFunc("/api/articles", api.GetArticleSummaries).Methods(http.MethodGet)
	r.HandleFunc("/api/articles/{article_id:[0-9]+}", api.GetArticle).Methods(http.MethodGet)
	r.HandleFunc("/api/articles", api.PostArticle).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.HandleFunc("/api/articles/{article_id:[0-9]+}/comments", api.GetArticleComments).Methods(http.MethodGet)
	r.HandleFunc("/api/articles/{article_id:[0-9]+}/comments", api.PostArticleComment).Methods(http.MethodPost).HeadersRegexp("Content-Type", "application/json")
	r.PathPrefix("/api/").HandlerFunc(api.NotFound)
	r.PathPrefix("/").Handler(asset).Methods(http.MethodGet)
	return handlers.LoggingHandler(os.Stdout, r)
}

func ErrorJSON(w http.ResponseWriter, message string, code int) {
	var out struct {
		Message string `json:"message"`
		Code    int    `json:"status_code"`
	}
	out.Message = message
	out.Code = code
	ResponseJSON(w, out, code)
}

func InternalServerError(w http.ResponseWriter) {
	ErrorJSON(w, "internal server error", http.StatusInternalServerError)
}

func ResponseJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
