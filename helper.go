package pregogen

import (
	"fmt"
	"go/ast"
	"log"
	"os"
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

// Get Array type
func GetArrayType(a string) string {
	return strings.Trim(a, "[]")
}

// Get Pointer type
func GetPointerType(a string) string {
	return strings.Trim(a, "*")
}

// IsArray checks if a field is an array.
func IsArray(fieldType string) bool {
	return strings.HasPrefix(fieldType, "[]")
}

// Seq creates a sequence of integers.
func Seq(start, size int) []int {
	res := make([]int, size)
	for i := range res {
		res[i] = i + start
	}
	return res
}

// Method gives action method specific translation in method (plus, append, or bytesBuffer)
// to avoid passing variables in subtemplate
func Method(method, action, addition string) (string, error) {
	switch method {
	case "bytesBuffer":
		return MethodBytesBuffer(action, addition)
	case "append":
		return MethodAppend(action, addition)
	case "plus":
		return MethodPlus(action, addition)
	case "unique":
		return MethodUnique(action, addition)
	}
	return "", fmt.Errorf("Unknown method: %v", method)
}

func MethodAppend(action, addition string) (string, error) {
	switch action {
	case "declareVar":
		return "\n\tvar result []byte\n\t", nil
	case "start":
		return "result = append(result, " + addition, nil
	case "stop":
		return addition + "...)\n\t", nil
	case "firststart":
		return "result = []byte(" + addition, nil
	case "firstend":
		return addition + ")\n\t", nil
	case "finalReturn":
		return "\n\treturn result, nil", nil
	}
	return "", fmt.Errorf("Unknown action: %v", action)
}

func MethodBytesBuffer(action, addition string) (string, error) {
	switch action {
	case "declareVar":
		return "\n\tvar result bytes.Buffer\n\t", nil
	case "start":
		return "result.WriteString(" + addition, nil
	case "stop":
		return addition + ")\n\t", nil
	case "firststart":
		return "result.WriteString(" + addition, nil
	case "firstend":
		return addition + ")\n\t", nil
	case "finalReturn":
		return "\n\treturn result.Bytes(), nil", nil
	}
	return "", fmt.Errorf("Unknown action: %v", action)
}

func MethodPlus(action, addition string) (string, error) {
	switch action {
	case "declareVar":
		return "\n\tvar result string\n\t", nil
	case "start":
		return "result += " + addition, nil // result += \","
	case "stop":
		return addition + "\n\t", nil
	case "firststart":
		return "result = " + addition, nil
	case "firstend":
		return addition + "\n\t", nil
	case "finalReturn":
		return "\n\treturn []byte(result), nil", nil
	}
	return "", fmt.Errorf("Unknown action: %v", action)
}

func MethodUnique(action, addition string) (string, error) {
	switch action {
	case "declareVar":
		return "", nil
	case "start":
		return "return ([]byte(", nil
	case "stop":
		return ")), nil", nil
	case "finalReturn":
		return "", nil
	}
	return "", fmt.Errorf("Unknown action: %v", action)
}

// Return.
func Return() error {
	os.Exit(0)
	return nil
}

// templateFuncs returns a FuncMap with the custom functions.
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		"dict":           dict,
		"sub":            Sub,
		"seq":            Seq,
		"getArrayType":   GetArrayType,
		"getPointerType": GetPointerType,
		"isArray":        IsArray,
		"method":         Method,
		"return":         Return,
	}
}

func processFieldType(fieldType, fieldName string, includes *[]string) {
	// Handle slice types
	fieldType = strings.TrimPrefix(fieldType, "[]")
	// Handle pointer types
	fieldType = strings.TrimPrefix(fieldType, "*")

	switch fieldType {
	case "string":
		// do nothing in this case
		log.Printf("Field %s is of type string", fieldName)
	case "bool":
		// do nothing in this case
		log.Printf("Field %s is of type bool", fieldName)
	case "int", "int16", "int32", "int64",
		"uint", "uint16", "uint32", "uint64",
		"float32", "float64":
		log.Printf("Field %s is of type number", fieldName)
		if !containsItem(*includes, "strconv") {
			*includes = append(*includes, "strconv")
		}
	case "time.Time":
		// no need to do anything in this case
		log.Printf("Field %s is of type time.Time", fieldName)
	case "byte", "int8", "uint8":
		// do nothing in this case
		log.Printf("Field %s is of type byte", fieldName)
		/*if !containsItem(*includes, "strconv") {
			*includes = append(*includes, "strconv")
		}*/
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
