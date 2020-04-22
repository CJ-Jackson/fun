package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CJ-Jackson/fun/commonUtil"
	"github.com/CJ-Jackson/fun/httpUtil"
	"github.com/CJ-Jackson/fun/prGen"
	"github.com/CJ-Jackson/fun/serveFile"
	"github.com/cjtoolkit/ctx"
)

func boot() (http.Handler, commonUtil.Param) {
	context := ctx.NewBackgroundContext()
	defer ctx.ClearBackgroundContext(context)

	param := commonUtil.GetParam(context)

	fmt.Println("Booting up application")

	serveFile.Boot(context)

	prGen.Boot(context)

	fmt.Println("Booted up successfully.")
	fmt.Println()

	return createHandler(httpUtil.GetRouter(context)), param
}

func createHandler(handler http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		header := res.Header()
		header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		header.Set("Service-Worker-Allowed", "/")
		req, _ = ctx.NewContext(res, req)
		handler.ServeHTTP(res, req)
	}
}

func main() {
	handler, param := boot()

	commonUtil.CheckIfTestRun(param)

	fmt.Println("Now listening on", param.Address)
	log.Print(http.ListenAndServe(param.Address, handler))
}
