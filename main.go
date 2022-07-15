package main

import (
	"github.com/wangyulu/web-go/app/console"
	"github.com/wangyulu/web-go/app/http"
	"github.com/wangyulu/web-go/app/provider/demo"
	"github.com/wangyulu/web-go/framework"
	"github.com/wangyulu/web-go/framework/provider/app"
	"github.com/wangyulu/web-go/framework/provider/distributed"
	"github.com/wangyulu/web-go/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewHadeContainer()

	// 绑定App服务提供者
	container.Bind(&app.HadeAppProvider{})

	// 后续初始化需要绑定的服务提供者...
	container.Bind(&demo.DemoProvider{})

	// 分布锁
	container.Bind(&distributed.LocalDistributedProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	// 运行root命令
	console.RunCommand(container)
}
