package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

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
	templeteFilePath := fileName + ".template"
	return packBox.FindString(templeteFilePath)
}

func generateFile(project *Project, fileNames []string) error {
	targetPath := project.BuildPath
	for _, fileName := range fileNames {
		targetFileName := filepath.Join(targetPath, fileName)
		buildPath := filepath.Dir(targetFileName)
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
		file, err := os.Create(targetFileName)
		if err != nil {
			return nil
		}
		err = tmpl.Execute(file, project)
		if err != nil {
			file.Close()
			return err
		}
		file.Close()

	}
	return nil
}
