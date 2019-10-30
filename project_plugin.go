package main

import (
	"github.com/gobuffalo/packr/v2"
)

var packBox *packr.Box

type templeteArgs struct {
	Name       string
	ModuleName string
}

func init() {
	packBox = GetPackBox()
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

//loadTemplete
func LoadTemplete(fileName string) (str string, err error) {
	templeteFilePath := fileName + ".template"
	return packBox.FindString(templeteFilePath)
}
