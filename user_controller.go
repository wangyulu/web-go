package main

import (
	"github.com/wangyulu/web-go/framework"
)

func UserLoginController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, UserLoginController")
	return nil
}
