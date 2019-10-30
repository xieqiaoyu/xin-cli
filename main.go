package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type Project struct {
	Name       string
	ModuleName string
	BuildPath  string
}

func testDir(targetPath string, shouldBeEmpty bool) error {
	f, err := os.Open(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s not exists ,%w", targetPath, err)
		}
		return err
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s not a dir", targetPath)
	}
	if !shouldBeEmpty {
		_, err = f.Readdirnames(1) // Or f.Readdir(1)
		if err != io.EOF {
			return fmt.Errorf("%s not an empty dir", targetPath)
		}
	}

	return nil
}

func loadTemplete(fileName string) (str string, err error) {
	templeteFilePath := "templetes/" + fileName + ".templete"
	tbytes, err := ioutil.ReadFile(templeteFilePath)
	return string(tbytes), err
}

func generateFile(project *Project, fileNames []string) error {
	targetPath := project.BuildPath
	for _, fileName := range fileNames {
		buildPath := filepath.Join(targetPath, filepath.Dir(fileName))
		err := os.MkdirAll(buildPath, os.ModePerm)
		if err != nil {
			return err
		}
		templeteString, err := loadTemplete(fileName)
		if err != nil {
			return err
		}
		tmpl, err := template.New(fileName).Parse(templeteString)
		if err != nil {
			return err
		}
		err = tmpl.Execute(os.Stdout, project)
		if err != nil {
			return err
		}

	}
	return nil
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
