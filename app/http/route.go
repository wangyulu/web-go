package http

import (
	"github.com/wangyulu/web-go/app/http/middleware/cors"
	"github.com/wangyulu/web-go/app/http/module/demo"
	"github.com/wangyulu/web-go/framework/contract"
	"github.com/wangyulu/web-go/framework/gin"
	"github.com/wangyulu/web-go/framework/middleware"
	ginSwagger "github.com/wangyulu/web-go/framework/middleware/gin-swagger"
	swaggerFiles "github.com/wangyulu/web-go/framework/middleware/gin-swagger/files"
	"github.com/wangyulu/web-go/framework/middleware/static"
)

func Routes(r *gin.Engine) {
	container := r.GetContainer()
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	r.Use(middleware.Trace())

	r.Use(cors.Default())

	// 如果配置了swagger，则显示swagger的中间件
	if configService.GetBool("app.swagger") == true {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	demo.Register(r)
}
