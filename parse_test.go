package pregogen

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func Test_parseSourceFile(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		wantErr  bool
	}{
		{
			name:     "Valid file",
			fileName: "tests/mix3types/mytypes.go",
			wantErr:  false,
		},
		{
			name:     "Non-existent file",
			fileName: "nonexistent.go",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSourceFile(tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSourceFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("parseSourceFile() = nil, want *ast.File")
			}
		})
	}
}

func Test_findVariableValue(t *testing.T) {
	// Create a simple AST for testing
	validFile := &ast.File{
		Name: &ast.Ident{Name: "test"},
		Decls: []ast.Decl{
			&ast.GenDecl{
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{{Name: "MyType_example"}},
						Values: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: "\"test value\"",
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name     string
		node     *ast.File
		typeName string
		wantErr  bool
	}{
		{
			name:     "Found variable",
			node:     validFile,
			typeName: "MyType",
			wantErr:  false,
		},
		{
			name:     "Variable not found",
			node:     validFile,
			typeName: "NonExistentType",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findVariableValue(tt.node, tt.typeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("findVariableValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("findVariableValue() = nil, want ast.Expr")
			}
		})
	}
}

func Test_extractValue(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want interface{}
	}{
		{
			name: "String literal",
			expr: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"test\"",
			},
			want: "test",
		},
		{
			name: "Integer literal",
			expr: &ast.BasicLit{
				Kind:  token.INT,
				Value: "123",
			},
			want: 123,
		},
		{
			name: "Float literal",
			expr: &ast.BasicLit{
				Kind:  token.FLOAT,
				Value: "12.34",
			},
			want: 12.34,
		},
		{
			name: "Boolean true",
			expr: &ast.Ident{
				Name: "true",
			},
			want: true,
		},
		{
			name: "Boolean false",
			expr: &ast.Ident{
				Name: "false",
			},
			want: false,
		},
		{
			name: "Struct literal",
			expr: &ast.CompositeLit{
				Elts: []ast.Expr{
					&ast.KeyValueExpr{
						Key: &ast.Ident{Name: "Field"},
						Value: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"value\"",
						},
					},
				},
			},
			want: map[string]interface{}{"Field": "value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractValue(tt.expr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
