package cmd

import (
	"fmt"
	"os"

	"lowframe/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var cfg = new(config.Cfg)

func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(cfg); err != nil {
		panic(fmt.Errorf("unmarshal conf server failed, err:%s ", err))
	}
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config related",
	Long:  `config related`,
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show config",
	Long:  `show config`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("server = %+v\n", cfg)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "settings.yaml", "config file")

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(showCmd)
}
