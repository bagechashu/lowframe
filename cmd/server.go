package cmd

import (
	"lowframe/service"

	"github.com/spf13/cobra"
)

var server = new(service.Server)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  `server`,
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "start server",
	Long:  `start server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(":" + cfg.Server.Port)
	},
}

var walkCmd = &cobra.Command{
	Use:   "walk",
	Short: "walk route server",
	Long:  `walk route server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.WalkRoutes()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(runCmd)
	serverCmd.AddCommand(walkCmd)
}
