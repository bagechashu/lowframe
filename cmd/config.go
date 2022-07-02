package cmd

import (
	"fmt"
	"os"

	"lowframe/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var serverCfg = new(config.Server)

func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(serverCfg); err != nil {
		panic(fmt.Errorf("unmarshal conf server failed, err:%s \n", err))
	}
}

var showConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "config",
	Long:  `config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("server = %+v\n", serverCfg)

	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "settings.yaml", "config file")

	rootCmd.AddCommand(showConfigCmd)
}
