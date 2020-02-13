package reqjsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"gopkg.in/yaml.v2"
	"strings"
)

const (
	SchemaCommentTag = "xin::yaml2jsonschema"
)

type Schemas struct {
	Package string
	Vars    map[string]string
}

func LoadAndParse() (*Schemas, error) {
	fset := token.NewFileSet() // positions are relative to fset
	// Only support current dir now
	d, err := parser.ParseDir(fset, "./", nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	var parseRes = make(map[string]string)
	var packageName string
	var parseOK = true

	for _pName, f := range d {
		// exclude test package
		if strings.HasSuffix(_pName, "_test") {
			continue
		}
		packageName = _pName
		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncDecl:
				var commentStr string
				commentStr = x.Doc.Text()
				schemaStartPos := strings.Index(commentStr, SchemaCommentTag)
				if schemaStartPos < 0 {
					break
				}
				schemaFirstLineEndPos := strings.Index(commentStr[schemaStartPos:], "\n")
				if schemaFirstLineEndPos < 0 {
					break
				}
				schemaDefEndPos := schemaStartPos + schemaFirstLineEndPos
				schemaDef := commentStr[(schemaStartPos + len(SchemaCommentTag)):schemaDefEndPos]
				schemaDefFields := strings.Fields(schemaDef)

				var schemaName, eofMark string
				schemaDefFieldLen := len(schemaDefFields)
				if schemaDefFieldLen > 0 {
					schemaName = schemaDefFields[0]
					if strings.HasPrefix(schemaName, "<<") {
						eofMark = schemaName[2:]
						if schemaDefFieldLen > 1 {
							schemaName = schemaDefFields[1]
						} else {
							// the first token is eof token, reset schemaName
							schemaName = ""
						}
					}
				}
				if schemaName == "" {
					//auto generate a schemaName
					schemaName = x.Name.String() + "RequestSchema"
				}

				yamlschema := commentStr[schemaDefEndPos:]
				if eofMark != "" {
					schemaEndPos := strings.LastIndex(yamlschema, eofMark)
					if schemaEndPos > 0 {
						yamlschema = yamlschema[:schemaEndPos]
					}
				}
				if yamlschema != "" {
					jsonStr, err := Yaml2Json([]byte(yamlschema))
					if err != nil {
						fmt.Printf("error: %s\n", err)
						parseOK = false
						return true
					}
					if _, exists := parseRes[schemaName]; exists {
						fmt.Printf("error: conflict var:%s\n", schemaName)
						parseOK = false
						return true
					}
					//TODO: use github.com/qri-io/jsonschema  to check error in valid jsonschema
					parseRes[schemaName] = string(jsonStr)

				}
			}

			return true
		})
	}
	if !parseOK {
		return nil, errors.New("parse error")
	}
	if len(parseRes) < 1 {
		return nil, errors.New("No schema found")
	}
	return &Schemas{
		Vars:    parseRes,
		Package: packageName,
	}, nil
}

func Yaml2Json(yamldata []byte) (jsondata []byte, err error) {
	m := map[interface{}]interface{}{}
	err = yaml.Unmarshal(yamldata, &m)
	if err != nil {
		return nil, err
	}
	jsonStruct, err := jsonConvert(m)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonStruct)
}

// encoding/json not support map[interface{}]interface{} need a convert
func jsonConvert(m interface{}) (interface{}, error) {
	switch v := m.(type) {
	case map[interface{}]interface{}:
		res := map[string]interface{}{}
		for k, v2 := range v {
			convertv, err := jsonConvert(v2)
			if err != nil {
				return nil, err
			}
			switch k2 := k.(type) {
			case string:
				res[k2] = convertv
			default:
				return nil, fmt.Errorf("unsupport map key type:%T", k)
			}
		}
		return res, nil
	/*
		case []interface{}:
			res := make([]interface{}, len(v))
			for i, v2 := range v {
				convertv, err := jsonConvert(v2)
				if err != nil {
					return nil, err
				}
				res[i] = convertv
			}
			return res, nil
	*/
	default:
		return m, nil
	}
}
