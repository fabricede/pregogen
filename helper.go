package pregogen

import (
	"fmt"
	"go/ast"
	"log"
	"strings"
	"text/template"
)

// dict creates a map from a list of pairs.
// This function is used in the template to create a map of field names and values.
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call: %v", values)
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings: %v", values)
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// Sub subtracts b from a.
func Sub(a, b int) int {
	return a - b
}

// templateFuncs returns a FuncMap with the custom functions.
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"dict": dict,
		"sub":  Sub,
	}
}

func processFieldType(fieldType, fieldName string, includes *[]string) {
	// Handle slice types
	fieldType = strings.TrimPrefix(fieldType, "[]")

	switch fieldType {
	case "string":
		// do nothing in this case
		log.Printf("Field %s is of type string", fieldName)
	case "bool":
		// do nothing in this case
		log.Printf("Field %s is of type bool", fieldName)
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		log.Printf("Field %s is of type int", fieldName)
		if !containsItem(*includes, "strconv") {
			*includes = append(*includes, "strconv")
		}
	case "byte":
		// do nothing in this case
		log.Printf("Field %s is of type byte", fieldName)
	default:
		// Handle other types
		log.Printf("Field %s is of other type", fieldName)
		if !containsItem(*includes, "fmt") {
			*includes = append(*includes, "fmt")
		}
	}
}

func parseJSONTag(tag string) (string, map[string]bool) {
	// Parse the JSON tag to extract the field name and options
	if tag == "" {
		return "", nil
	}
	tag = strings.Trim(tag, "`")
	parts := strings.Split(tag, " ")
	tagMap := make(map[string]bool)
	for _, part := range parts {
		keyValue := strings.Split(part, ":")
		if len(keyValue) == 2 && keyValue[0] == "json" {
			jsonTag := strings.Trim(keyValue[1], "\"")
			tagParts := strings.Split(jsonTag, ",")
			fieldName := tagParts[0]
			for _, option := range tagParts[1:] {
				tagMap[option] = true
			}
			return fieldName, tagMap
		}
	}
	return "", tagMap
}

func getTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.ArrayType:
		return "[]" + getTypeName(t.Elt)
	case *ast.StarExpr:
		return "*" + getTypeName(t.X)
	case *ast.SelectorExpr:
		return getTypeName(t.X) + "." + t.Sel.Name
	default:
		return ""
	}
}

func containsItem(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
