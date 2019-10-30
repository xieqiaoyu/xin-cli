package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
)

var packBox *packr.Box

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

	packBox = packr.New("tBox", "./templates")

	project := &Project{
		Name:       "midas",
		ModuleName: "github.com/xieqiaoyu/midas",
		BuildPath:  "./project",
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
