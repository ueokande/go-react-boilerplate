package web

import (
	"net/http"
)

func (h APIHandler) Health(w http.ResponseWriter, r *http.Request) {
	ErrorJSON(w, "ok", http.StatusOK)
}

func (h APIHandler) Ready(w http.ResponseWriter, r *http.Request) {
	ErrorJSON(w, "ok", http.StatusOK)
}
