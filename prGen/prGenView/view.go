package prGenView

import (
	"fun/commonUtil"
	"fun/prGen/internal"
	"html/template"

	"github.com/cjtoolkit/ctx"
)

type PrGenView interface {
	ExecIndexView(context ctx.Context)
}

func GetPrGenView(context ctx.BackgroundContext) PrGenView {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return initPrGenView(), nil
	}).(PrGenView)
}

func initPrGenView() PrGenView {
	return pRGenView{
		indexHtml: template.Must(template.New("Index").Parse(commonUtil.DecodeValueStr(internal.Index))),
	}
}

type pRGenView struct {
	indexHtml *template.Template
}

func (p pRGenView) ExecIndexView(context ctx.Context) {
	p.indexHtml.Execute(context.ResponseWriter(), nil)
}
