package prGen

import (
	"fun/prGen/prGenController"
	"fun/urlMap"
	"net/http"

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

	b.router.POST(urlMap.Index, func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		context := ctx.GetContext(req)
		b.controller.PostIndex(context)
	})
}
