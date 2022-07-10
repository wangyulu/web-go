package main

import (
	"github.com/wangyulu/web-go/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)

	c.SetOkStatus().Json("ok, UserLoginController")
	return nil
}
