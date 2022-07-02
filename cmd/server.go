package cmd

import (
	"lowframe/service"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  `server`,
	Run: func(cmd *cobra.Command, args []string) {
		server := service.NewServer()
		server.Run(":" + serverCfg.Port)
	},
}
