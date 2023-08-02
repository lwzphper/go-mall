package cmd

import (
	"fmt"
	"github.com/lwzphper/go-mall/version"
	"github.com/spf13/cobra"
	"os"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

var showVersion bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-mall",
	Short: "go-mall",
	Long:  "go-mall",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			fmt.Print(version.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/config.toml", "the service config from file")
	RootCmd.PersistentFlags().StringVarP(&confETCD, "config-etcd", "e", "127.0.0.1:2379", "the service config from etcd")
	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "the go-mall version")
}
