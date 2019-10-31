package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "xin",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("need a sub command")
		},
	}
)

//RootCmd return root commend
func RootCmd() *cobra.Command {
	return rootCmd
}

func Execute() {
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(CreateCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}