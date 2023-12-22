package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var rootCmd = &cobra.Command{
	Use:   "mydemoapp",
	Short: "演示程序",
	Long:  "演示delve使用方法所做的程序",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
