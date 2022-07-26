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
	// cmd
	root.AddCommand(initCmdCommand())
	// provider
	root.AddCommand(initProviderCommand())
	// middleware
	root.AddCommand(initMiddlewareCommand())
	// new
	root.AddCommand(initNewCommand())
	// swagger
	root.AddCommand(initSwaggerCommand())
}
