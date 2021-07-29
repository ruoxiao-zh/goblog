package main

import (
	"net/http"

	"github.com/ruoxiao-zh/goblog/app/http/middlewares"
	"github.com/ruoxiao-zh/goblog/bootstrap"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
