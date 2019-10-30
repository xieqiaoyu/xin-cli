package main

import (
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"os"
)

var packBox *packr.Box

type Project struct {
	Name       string
	ModuleName string
	BuildPath  string
}

func init() {
	packBox = packr.New("tBox", "./templates")
}
func main() {
	project := &Project{
		Name:       "midas",
		ModuleName: "github.com/xieqiaoyu/midas",
		BuildPath:  "./project",
	}
	err := testDir(project.BuildPath, true)
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
