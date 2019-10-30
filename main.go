package main

import (
	"fmt"
	"os"
)

func main() {
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
