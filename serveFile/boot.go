package serveFile

import (
	"fun/commonUtil"
	"fun/httpUtil"
	"fun/urlMap"
	"net/http"
	"os"

	"github.com/cjtoolkit/ctx"
	"github.com/cjtoolkit/zipfs"
	"github.com/julienschmidt/httprouter"
)

func Boot(context ctx.BackgroundContext) {
	router := httpUtil.GetRouter(context)

	if commonUtil.GetParam(context).Production == false {
		if _, err := os.Stat("asset/live"); err == nil {
			serveFiles(router, http.Dir("asset/live"))
			return
		}
	}
	serveFiles(router, zipfs.InitZipFs("asset.zip"))
}

func serveFiles(router *httprouter.Router, fs http.FileSystem) {
	router.ServeFiles(urlMap.FontsFiles, zipfs.Prefix("/fonts", fs))
	router.ServeFiles(urlMap.JavascriptsFiles, zipfs.Prefix("/javascripts", fs))
	router.ServeFiles(urlMap.StylesheetFiles, zipfs.Prefix("/stylesheets", fs))
}
