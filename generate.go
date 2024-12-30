// Package main implements a code generator for adding custom JSON encoding and decoding methods
// to Go structs without using the standard json library. This generator uses go generate to
// parse Go source files, identify structs, and generate the necessary methods based on a template.
package pregogen

import (
	"embed"
	"fmt"
	"go/ast"
	"log"
	"os"
	"strings"
	"text/template"
)

//go:embed json_templates/*.tmpl
var templatesFS embed.FS

// FieldInfo holds information about a struct field, including its name, JSON tag name, and whether it is a string.
type FieldInfo struct {
	Name         string
	JSONName     string
	Properties   map[string]bool
	FieldType    string
	ReceiverName string
}

// TemplateData holds the data required to populate the template for generating JSON methods,
// including the package name, type name, receiver name, and a list of fields.
type TemplateData struct {
	Package      string
	TypeName     string
	Fields       []FieldInfo
	ReceiverName string
	Includes     []string
	Gentype      string
}

// jsonData holds the JSON data of the example value for tests.
var jsonData []byte

func RunGenerator(typeName, fileName, genType string) ([]byte, error) {
	node, err := parseSourceFile(fileName)
	if err != nil {
		return nil, err
	}
	value, err := findVariableValue(node, typeName)
	if err != nil {
		return nil, err
	}
	jsonData, err = marshalExampleValue(value)
	if err != nil {
		return nil, err
	}

	err = generateForStruct(node, typeName, genType)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func generateForStruct(node *ast.File, typeName, genType string) error {
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Name.Name == typeName {
				generateJSONMarshal(x, node.Name.Name, genType)
				return false
			}
		}
		return true
	})
	return nil
}

func generateJSONMarshal(typeSpec *ast.TypeSpec, packageName string, gentype string) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		log.Fatalf("%s is not a struct", typeSpec.Name.Name)
	}

	receiverName := strings.ToLower(string(typeSpec.Name.Name[0]))
	var includes []string
	var fields []FieldInfo
	for _, field := range structType.Fields.List {
		jsonTag := ""
		if field.Tag != nil {
			jsonTag = field.Tag.Value
		}
		for _, name := range field.Names {
			fieldType := getTypeName(field.Type)
			jsonName, parsedTags := parseJSONTag(jsonTag)
			fields = append(fields, FieldInfo{
				Name:         name.Name,
				JSONName:     jsonName,
				Properties:   parsedTags,
				FieldType:    fieldType,
				ReceiverName: receiverName,
			})
			processFieldType(fieldType, name.Name, &includes)
		}
	}

	for i, inc := range includes {
		log.Printf("Include: %d,%s", i, inc)
	}

	data := TemplateData{
		Package:      packageName,
		TypeName:     typeSpec.Name.Name,
		ReceiverName: receiverName,
		Fields:       fields,
		Includes:     includes,
		Gentype:      gentype,
	}

	tmpl, err := template.New("marshal.tmpl").Funcs(templateFuncs()).ParseFS(templatesFS, "json_templates/*.tmpl")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	file, err := os.Create(fmt.Sprintf("generated_%s_marshal_json_%s.go", typeSpec.Name.Name, gentype))
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	err = tmpl.ExecuteTemplate(file, "marshal.tmpl", data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
