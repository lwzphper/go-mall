package cmd

import "github.com/spf13/cobra"

// 项目初始化命令

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "go-mail 初始化",
	Long:  "go-mail 初始化",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局配置
		if err := loadGlobalConfig(); err != nil {
			return err
		}

		// 创建数据表
		err := createTables()
		if err != nil {
			return err
		}

		return nil
	},
}

// 初始化全局配置
func loadGlobalConfig() error {
	return nil
}

// 初始化数据表
func createTables() error {
	return nil
}

func init() {
	RootCmd.AddCommand(initCmd)
}
