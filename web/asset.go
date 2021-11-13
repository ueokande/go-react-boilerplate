package web

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type assetHandler struct {
	webroot string
}

func (h *assetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := path.Clean(r.URL.Path)
	p := filepath.Join(h.webroot, path.Clean(r.URL.Path))
	if strings.HasPrefix(req, "/static") {
		if info, err := os.Stat(p); err != nil || info.IsDir() {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, p)
		return
	}

	if info, err := os.Stat(p); err == nil && !info.IsDir() {
		http.ServeFile(w, r, p)
		return
	}
	http.ServeFile(w, r, filepath.Join(h.webroot, "index.html"))
}

func NewAssetHandler(webroot string) http.Handler {
	return &assetHandler{webroot: webroot}
}

func NewDebuAssetHandler(webpack *url.URL) http.Handler {
	return httputil.NewSingleHostReverseProxy(webpack)
}
