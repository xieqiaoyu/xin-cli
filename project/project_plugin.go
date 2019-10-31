package project

import (
	"github.com/gobuffalo/packr/v2"
)

var packBox *packr.Box

type templeteArgs struct {
	Name       string
	ModuleName string
}

func init() {
	packBox = packr.New("tBox", "./templates")
}

//GetBuildFiles 获取要构建的文件名列表（相对于构建路径）
func GetBuildFiles() ([]string, error) {
	return []string{
		"main.go",
		"go.mod",
		"config.toml",
		".gitignore",
		"cmd/playground.go",
		"cmd/server.go",
		"metadata/metadata.go",
		"router/router.go",
		"handle/restful/helloworld.go",
		//"static/configschema.json",
	}, nil
}

//GetTempleteArgs 获取用于模板渲染的参数对象
func GetTempleteArgs() interface{} {
	return &templeteArgs{
		Name:       "midas",
		ModuleName: "github.com/xieqiaoyu/midas",
	}
}

//LoadTemplete 给定文件名返回文件的模板
func LoadTemplete(fileName string) (str string, err error) {
	templeteFilePath := fileName + ".template"
	return packBox.FindString(templeteFilePath)
}
