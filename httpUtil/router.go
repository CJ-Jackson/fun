package httpUtil

import (
	"github.com/cjtoolkit/ctx"
	"github.com/julienschmidt/httprouter"
)

func GetRouter(context ctx.BackgroundContext) *httprouter.Router {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		router := httprouter.New()

		return router, nil
	}).(*httprouter.Router)
}
