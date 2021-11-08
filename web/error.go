package web

import (
	"net/http"
)

func (h APIHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	ErrorJSON(w, "not found", http.StatusNotFound)
}
