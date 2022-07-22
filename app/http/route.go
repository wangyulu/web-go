package http

import (
	"github.com/wangyulu/web-go/app/http/module/demo"
	"github.com/wangyulu/web-go/framework/gin"
	"github.com/wangyulu/web-go/framework/middleware"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	r.Use(middleware.Trace())

	demo.Register(r)
}
