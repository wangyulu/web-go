package main

import (
	"github.com/wangyulu/web-go/framework"
)

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}