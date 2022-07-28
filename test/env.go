package test

import (
	"github.com/wangyulu/web-go/framework"
	"github.com/wangyulu/web-go/framework/provider/app"
	"github.com/wangyulu/web-go/framework/provider/env"
)

const (
	BasePath = "/Users/sky/Code/local/gomodule/geek/web-go/"
)

func InitBaseContainer() framework.Container {
	// 初始化服务容器
	container := framework.NewHadeContainer()

	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{BaseFolder: BasePath})

	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.HadeTestingEnvProvider{})

	return container
}
