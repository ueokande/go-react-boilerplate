package web

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

type assetHandler struct {
	webroot string
}

func (h *assetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join(h.webroot, path.Clean(r.URL.Path))
	if info, err := os.Stat(p); err != nil || info.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.webroot, "index.html"))
		return
	}
	http.ServeFile(w, r, p)
}

func NewAssetHandler(webroot string) http.Handler {
	return &assetHandler{webroot: webroot}
}

func NewDebuAssetHandler(webpack *url.URL) http.Handler {
	return httputil.NewSingleHostReverseProxy(webpack)
}
