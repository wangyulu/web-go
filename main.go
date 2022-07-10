package main

import (
	"github.com/wangyulu/web-go/framework"
	"github.com/wangyulu/web-go/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()

	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)

	server := http.Server{
		// 自定义的请求核心处理函数
		Handler: core,

		Addr: ":8080",
	}

	server.ListenAndServe()
}
