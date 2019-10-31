package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xieqiaoyu/xin-cli/metadata"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "print the version number",
		Long:  `we also have a version number`,
		Run: func(cmd *cobra.Command, args []string) {
			metadata.ShowVersion()
		},
	}
}
