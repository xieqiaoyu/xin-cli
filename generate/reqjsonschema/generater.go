package reqjsonschema

import (
	"github.com/gobuffalo/packr/v2"
	"os"
	"text/template"
)

var packBox *packr.Box

const genfileName = "reqJsonSchema_gen.go"

func init() {
	packBox = packr.New("reqBox", "./templates")
}

func LoadTemplate(fileName string) (str string, err error) {
	templeteFilePath := fileName + ".template"
	return packBox.FindString(templeteFilePath)
}

func GenerateFile(schemas *Schemas) error {
	err := CompareOldFile(schemas, genfileName)
	if err != nil {
		println(err.Error())
	}

	templeteString, err := LoadTemplate(genfileName)
	if err != nil {
		return err
	}
	tmpl, err := template.New(genfileName).Parse(templeteString)
	if err != nil {
		return err
	}
	file, err := os.Create(genfileName)
	if err != nil {
		return err
	}
	err = tmpl.Execute(file, schemas)
	if err != nil {
		file.Close()
		return err
	}
	file.Close()
	return nil
}
