package prGenView

import (
	"html/template"
	template2 "text/template"

	"github.com/CJ-Jackson/fun/commonUtil"
	"github.com/CJ-Jackson/fun/prGen/internal"

	"github.com/cjtoolkit/ctx"
)

type PrGenView interface {
	ExecIndexView(context ctx.Context)
	ExecManifest(context ctx.Context)
}

func GetPrGenView(context ctx.BackgroundContext) PrGenView {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return initPrGenView(), nil
	}).(PrGenView)
}

func initPrGenView() PrGenView {
	return pRGenView{
		indexHtml:    template.Must(template.New("Index").Parse(commonUtil.DecodeValueStr(internal.Index))),
		prTxtTpl:     template2.Must(template2.New("PR").Parse(commonUtil.DecodeValueStr(internal.Pr))),
		manifestJson: commonUtil.DecodeValue(internal.Manifest),
	}
}

type pRGenView struct {
	indexHtml    *template.Template
	prTxtTpl     *template2.Template
	manifestJson []byte
}

func (p pRGenView) ExecIndexView(context ctx.Context) {
	p.indexHtml.Execute(context.ResponseWriter(), nil)
}

func (p pRGenView) ExecManifest(context ctx.Context) {
	res := context.ResponseWriter()
	res.Header().Set("Content-Type", "application/json")
	res.Write(p.manifestJson)
}
