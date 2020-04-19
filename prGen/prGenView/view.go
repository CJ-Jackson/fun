package prGenView

import (
	"fun/commonUtil"
	"fun/prGen/internal"
	"fun/prGen/prGenModel"
	"html/template"
	template2 "text/template"

	"github.com/cjtoolkit/ctx"
)

type PrGenView interface {
	ExecIndexView(context ctx.Context)
	ExecPostIndexView(context ctx.Context, model prGenModel.PrGenModel)
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
		prTxtTpl:  template2.Must(template2.New("PR").Parse(commonUtil.DecodeValueStr(internal.Pr))),
	}
}

type pRGenView struct {
	indexHtml *template.Template
	prTxtTpl  *template2.Template
}

func (p pRGenView) ExecIndexView(context ctx.Context) {
	p.indexHtml.Execute(context.ResponseWriter(), nil)
}

func (p pRGenView) ExecPostIndexView(context ctx.Context, model prGenModel.PrGenModel) {
	p.prTxtTpl.Execute(context.ResponseWriter(), model)
}
