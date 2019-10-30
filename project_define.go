package main

import (
	"github.com/gobuffalo/packr/v2"
)

type templeteArgs struct {
	Name       string
	ModuleName string
}

func GetBuildFiles() ([]string, error) {
	return []string{
		"main.go",
		".gitignore",
		"cmd/playground.go",
		"metadata/metadata.go",
	}, nil
}

func GetTempleteArgs() interface{} {
	return &templeteArgs{
		Name:       "midas",
		ModuleName: "github.com/xieqiaoyu/midas",
	}
}

func GetPackBox() *packr.Box {
	return packr.New("tBox", "./templates")
}
