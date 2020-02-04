package project

import (
	"fmt"
	"github.com/c-bata/go-prompt"
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
		"go.mod",
		"makefile",
		"run-dev.sh",
		"metadata/metadata.go",
		"http/service.go",
		"http/handler/restful/handler.go",
		"http/handler/restful/helloworld.go",
		"cmd/main.go",
		"cmd/wire.go",
		"cmd/inject_cmd.go",
		"cmd/inject_http.go",
		"config.toml",
		".gitignore",
	}, nil
}
func completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

//GetTempleteArgs 获取用于模板渲染的参数对象
func GetTempleteArgs() interface{} {
	var projectName, moduleName string
	for projectName == "" {
		fmt.Println("the project name ?")
		projectName = prompt.Input("> ", completer)
	}
	for moduleName == "" {
		fmt.Println("the module name ?")
		moduleName = prompt.Input("> ", completer)
	}

	return &templeteArgs{
		Name:       projectName,
		ModuleName: moduleName,
	}
}

//LoadTemplete 给定文件名返回文件的模板
func LoadTemplete(fileName string) (str string, err error) {
	templeteFilePath := fileName + ".template"
	return packBox.FindString(templeteFilePath)
}
