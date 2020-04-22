package prGen

import (
	"github.com/CJ-Jackson/fun/httpUtil"
	"github.com/CJ-Jackson/fun/prGen/prGenController"

	"github.com/cjtoolkit/ctx"
)

func Boot(context ctx.BackgroundContext) {
	controllerBootKit{
		controller: prGenController.NewPrGenController(context),
		router:     httpUtil.GetRouter(context),
	}.boot()
}
