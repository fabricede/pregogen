package pregogen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

func parseSourceFile(fileName string) (*ast.File, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func findVariableValue(node *ast.File, typeName string) (ast.Expr, error) {
	varName := typeName + "_example"
	var value ast.Expr
	ast.Inspect(node, func(n ast.Node) bool {
		if decl, ok := n.(*ast.ValueSpec); ok {
			for i, name := range decl.Names {
				if name.Name == varName {
					if i < len(decl.Values) {
						value = decl.Values[i]
						return false
					}
				}
			}
		}
		return true
	})
	if value == nil {
		return nil, fmt.Errorf("variable %s not found", varName)
	}
	return value, nil
}

func extractValue(expr ast.Expr) interface{} {
	switch v := expr.(type) {
	case *ast.CompositeLit:
		// Handle struct literal
		result := make(map[string]interface{})
		for _, elt := range v.Elts {
			if kv, ok := elt.(*ast.KeyValueExpr); ok {
				key := kv.Key.(*ast.Ident).Name
				result[key] = extractValue(kv.Value)
			}
		}
		return result
	case *ast.BasicLit:
		// Handle basic literals (strings, numbers, etc.)
		switch v.Kind {
		case token.INT:
			val, _ := strconv.Atoi(v.Value)
			return val
		case token.FLOAT:
			val, _ := strconv.ParseFloat(v.Value, 64)
			return val
		case token.STRING:
			return strings.Trim(v.Value, "\"")
		}
	case *ast.Ident:
		// Handle identifiers (true, false)
		if v.Name == "true" {
			return true
		}
		if v.Name == "false" {
			return false
		}
	}
	return nil
}
