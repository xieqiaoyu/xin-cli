package main

import (
	"fmt"
)

type Project struct {
	Name       string
	ModuleName string
	BuildPath  string
}

func main() {

	var allfiles = []string{
		"main.go",
		".gitignore",
		"cmd/playground.go",
		"metadata/metadata.go",
	}
	project := &Project{
		Name:       "midas",
		ModuleName: "github.com/xieqiaoyu/midas",
		BuildPath:  "./artifact",
	}
	err := testDir(project.BuildPath, true)
	if err != nil {
		fmt.Println(err)
	}
	err = generateFile(project, allfiles)
	if err != nil {
		fmt.Println(err)
	}
}
