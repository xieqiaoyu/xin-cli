package project

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

//TestDir  test target path is avaliable ,dir with be maked if not exist
func TestDir(targetPath string, shouldBeEmpty bool) error {
	f, err := os.Open(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}
	defer f.Close()
	fileInfo, err := f.Stat()
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s not a dir", targetPath)
	}
	if shouldBeEmpty {
		_, err = f.Readdirnames(1) // Or f.Readdir(1)
		if err != io.EOF {
			return fmt.Errorf("%s not an empty dir", targetPath)
		}
	}

	return nil
}

func GenerateFile(project *Project, fileNames []string) error {
	targetPath := project.BuildPath
	for _, fileName := range fileNames {
		targetFileName := filepath.Join(targetPath, fileName)
		buildPath := filepath.Dir(targetFileName)
		err := os.MkdirAll(buildPath, os.ModePerm)
		if err != nil {
			return err
		}
		templeteString, err := LoadTemplete(fileName)
		if err != nil {
			return err
		}
		tmpl, err := template.New(fileName).Parse(templeteString)
		if err != nil {
			return err
		}
		file, err := os.Create(targetFileName)
		if err != nil {
			return nil
		}
		err = tmpl.Execute(file, project.TArgs)
		if err != nil {
			file.Close()
			return err
		}
		file.Close()
	}
	return nil
}
