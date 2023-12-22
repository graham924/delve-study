package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "nobody", "user name")
	createCmd.MarkFlagRequired("name")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "子命令",
	Long:  "做一个子命令 - create",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create command is called")
		name, _ := cmd.Flags().GetString("name")
		create(name)
	},
}

func create(name string) {
	fmt.Println("name: ", name)
}
