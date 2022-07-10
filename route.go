package main

import (
	"github.com/wangyulu/web-go/framework/gin"
	"github.com/wangyulu/web-go/framework/middleware"
)

func registerRouter(core *gin.Engine) {
	// 静态路由+HTTP方法匹配
	core.GET("/user/login", UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test3())

		// 动态路由
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", middleware.Test1(), SubjectGetController)
		subjectApi.GET("/list/all", middleware.Test2(), SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}
