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

func RunGenerator(typeName, fileName, genType string) error {
	node, err := parseSourceFile(fileName)
	if err != nil {
		return err
	}
	err = generateForStruct(node, typeName, genType)
	return err
}

func generateForStruct(node *ast.File, typeName, genType string) error {
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Name.Name == typeName {
				generateJSON(x, node.Name.Name, genType)
				return false
			}
		}
		return true
	})
	return nil
}

func generateJSON(typeSpec *ast.TypeSpec, packageName string, gentype string) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		log.Fatalf("%s is not a struct", typeSpec.Name.Name)
	}

	receiverName := strings.ToLower(string(typeSpec.Name.Name[0])) + "_receiver"
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

	TypeName := typeSpec.Name.Name
	name := fmt.Sprintf("generated_%s_marshal_json_%s.go", TypeName, gentype)
	if gentype == "unmarshal" {
		name = fmt.Sprintf("generated_%s_unmarshal_json.go", typeSpec.Name.Name)
	} else if gentype == "testMarshal" || gentype == "testUnmarshal" || gentype == "testAll" {
		name = fmt.Sprintf("generated_%s_json_test.go", typeSpec.Name.Name)
	} else if gentype == "marshal" {
		for _, method := range GetMethods() {
			name = fmt.Sprintf("generated_%s_marshal_json_%s.go", TypeName, method)
			processGentype(method, packageName, receiverName, TypeName, name, fields, includes)
		}
		return
	}
	processGentype(gentype, packageName, receiverName, TypeName, name, fields, includes)
}

func processGentype(gentype, packageName, receiverName, TypeName, filename string, fields []FieldInfo, includes []string) {
	for i, inc := range includes {
		log.Printf("Include: %d,%s", i, inc)
	}

	data := TemplateData{
		Package:      packageName,
		TypeName:     TypeName,
		ReceiverName: receiverName,
		Fields:       fields,
		Includes:     includes,
		Gentype:      gentype,
	}

	tmpl, err := template.New("json.tmpl").Funcs(templateFuncs()).ParseFS(templatesFS, "json_templates/*.tmpl")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	err = tmpl.ExecuteTemplate(file, "json.tmpl", data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
