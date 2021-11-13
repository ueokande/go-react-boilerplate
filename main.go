package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"

	"github.com/ueokande/go-react-builderplate/repository/mock"
	"github.com/ueokande/go-react-builderplate/web"
)

var flgDebug = flag.Bool("debug", false, "Run as a debug mode")
var flgHttp = flag.String("http", ":8000", "Host and port the server listen on")
var flgWebpack = flag.String("webpack", "http://127.0.0.1:3000", "Upstream webpack server for debug mode.  If debug mode is disabled, the assets are provided from local filesystem")
var flgWebroot = flag.String("webroot", "./frontend/build", "Path to asset root directory for production mode.  If debug mode is enabled, the assets are provided via webpack debug server")

func main() {
	flag.Parse()

	var asset http.Handler
	if *flgDebug {
		u, err := url.Parse(*flgWebpack)
		if err != nil {
			log.Fatal(err)
		}
		asset = web.NewDebuAssetHandler(u)
	} else {
		asset = web.NewAssetHandler(*flgWebroot)
	}
	repo := mock.Repository{}
	api := web.APIHandler{
		ArticleRepository: repo,
		CommentRepository: repo,
	}
	w := web.NewWeb(asset, api)

	log.Printf("Starting server on %s", *flgHttp)
	err := http.ListenAndServe(*flgHttp, w)
	if err != nil {
		log.Fatal(err)
	}
}
