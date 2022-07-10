package main

import (
	"github.com/wangyulu/web-go/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)

	c.ISetOkStatus().IJson("ok, UserLoginController")
}
