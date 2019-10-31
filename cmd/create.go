package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xieqiaoyu/xin-cli/project"
	"os"
)

func CreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create [BUILDPATH]",
		Short: "create a new project",
		Long:  `create a new project`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 0 {
				fmt.Println("need a build path")
				os.Exit(1)
			}
			buildpath := args[0]
			err := project.TestDir(buildpath, true)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			templeteArgs := project.GetTempleteArgs()
			newProject := &project.Project{
				BuildPath: buildpath, //TODO:命令行指定
				TArgs:     templeteArgs,
			}

			allfiles, err := project.GetBuildFiles()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = project.GenerateFile(newProject, allfiles)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
}
