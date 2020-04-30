package prGen

import (
	"net/http"

	"github.com/CJ-Jackson/fun/prGen/prGenController"
	"github.com/CJ-Jackson/fun/urlMap"
	"github.com/cjtoolkit/ctx"
	"github.com/julienschmidt/httprouter"
)

type controllerBootKit struct {
	controller prGenController.PrGenController
	router     *httprouter.Router
}

func (b controllerBootKit) boot() {
	b.router.GET(urlMap.Index, func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		context := ctx.GetContext(req)
		b.controller.Index(context)
	})

	b.router.GET(urlMap.Manifest, func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		context := ctx.GetContext(req)
		b.controller.Manifest(context)
	})
}
