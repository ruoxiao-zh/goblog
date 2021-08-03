package main

import (
	"net/http"

	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	c "goblog/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()

	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	_ = http.ListenAndServe(":" + c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
