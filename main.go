package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"os"
)

var packBox *packr.Box

func main() {
	packBox = GetPackBox()
	templeteArgs := GetTempleteArgs()
	project := &Project{
		BuildPath: "./project", //TODO:命令行指定
		tArgs:     templeteArgs,
	}

	allfiles, err := GetBuildFiles()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = testDir(project.BuildPath, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = generateFile(project, allfiles)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
