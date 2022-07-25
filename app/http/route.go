package http

import (
	"github.com/wangyulu/web-go/app/http/middleware/cors"
	"github.com/wangyulu/web-go/app/http/module/demo"
	"github.com/wangyulu/web-go/framework/gin"
	"github.com/wangyulu/web-go/framework/middleware"
	"github.com/wangyulu/web-go/framework/middleware/static"
)

func Routes(r *gin.Engine) {
	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	r.Use(middleware.Trace())

	r.Use(cors.Default())

	demo.Register(r)
}
