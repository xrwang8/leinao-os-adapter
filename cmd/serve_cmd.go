package cmd

import (
	"github.com/spf13/cobra"
	_ "leinao-os-adapter/config"
	"leinao-os-adapter/pkg/global/variable"
	"leinao-os-adapter/routers"
)

var serveCmd = &cobra.Command{
	Use:          "serve",
	Short:        "start leinao-os-adapter",
	SilenceUsage: true,
	Long:         "leinao-os-adapterStart service",
	Run: func(c *cobra.Command, args []string) {
		router := routers.InitWebRouter()
		_ = router.Run(variable.ConfigYml.GetString("HttpServer.Web.Port"))
	},
}
