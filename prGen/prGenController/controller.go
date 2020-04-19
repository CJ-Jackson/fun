package prGenController

import (
	"encoding/json"
	"fun/prGen/prGenModel"
	"fun/prGen/prGenView"

	"github.com/cjtoolkit/ctx"
)

type PrGenController struct {
	prGenView prGenView.PrGenView
}

func NewPrGenController(context ctx.BackgroundContext) PrGenController {
	return PrGenController{
		prGenView: prGenView.GetPrGenView(context),
	}
}

func (c PrGenController) Index(context ctx.Context) {
	c.prGenView.ExecIndexView(context)
}

func (c PrGenController) PostIndex(context ctx.Context) {
	model := prGenModel.PrGenModel{}
	json.NewDecoder(context.Request().Body).Decode(&model)
	c.prGenView.ExecPostIndexView(context, model)
}
