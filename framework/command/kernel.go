package command

import (
	"github.com/wangyulu/web-go/framework/cobra"
)

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCommand)

	// app
	root.AddCommand(initAppCommand())
	// env
	root.AddCommand(envCommand)
	// cron
	root.AddCommand(initCronCommand())
	// config
	root.AddCommand(initConfigCommand())
	// build
	root.AddCommand(initBuildCommand())
	// go
	root.AddCommand(goCommand)
	// npm
	root.AddCommand(npmCommand)
	// dev
	root.AddCommand(initDevCommand())

	// root.AddCommand(deployCommand)
	//
	// // cmd
	// cmdCommand.AddCommand(cmdListCommand)
	// cmdCommand.AddCommand(cmdCreateCommand)
	// root.AddCommand(cmdCommand)
	//
	//
	// // middleware
	// middlewareCommand.AddCommand(middlewareAllCommand)
	// middlewareCommand.AddCommand(middlewareAddCommand)
	// middlewareCommand.AddCommand(middlewareRemoveCommand)
	// root.AddCommand(middlewareCommand)
	//
	// // swagger
	// swagger.IndexCommand.AddCommand(swagger.InitServeCommand())
	// swagger.IndexCommand.AddCommand(swagger.GenCommand)
	// root.AddCommand(swagger.IndexCommand)
	//
	// // provider
	// providerCommand.AddCommand(providerListCommand)
	// providerCommand.AddCommand(providerCreateCommand)
	// root.AddCommand(providerCommand)
	//
	// // new
	// root.AddCommand(initNewCommand())
}
