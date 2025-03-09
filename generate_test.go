package pregogen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func Test_RunGenerator(t *testing.T) {
	// Create a temporary test directory and file
	testContent := `package testgen
type TestStruct struct {
	Field string ` + "`json:\"field\"`" + `
}
`
	tmpParent := filepath.Join(os.TempDir(), "pregogen_test")
	tmpdir := filepath.Join(tmpParent, "src", "testgen") // Create a proper package structure
	err := os.MkdirAll(tmpdir, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpParent)

	testFile := filepath.Join(tmpdir, "test.go")
	if err := os.WriteFile(testFile, []byte(testContent), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		typeName string
		fileName string
		genType  string
		wantErr  bool
	}{
		{
			name:     "Valid generation",
			typeName: "TestStruct",
			fileName: testFile,
			genType:  "marshal",
			wantErr:  false,
		},
		{
			name:     "Invalid file",
			typeName: "TestStruct",
			fileName: "nonexistent.go",
			genType:  "marshal",
			wantErr:  true,
		},
		{
			name:     "Invalid type",
			typeName: "NonExistentType",
			fileName: testFile,
			genType:  "marshal",
			wantErr:  false, // No error since it just skips non-existent types
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RunGenerator(tt.typeName, tt.fileName, tt.genType)
			if (err != nil) != tt.wantErr {
				t.Errorf("RunGenerator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && tt.typeName == "TestStruct" {
				// Verify files were generated in the correct directory
				for _, method := range GetMethods() {
					expectedFile := filepath.Join(tmpdir, fmt.Sprintf("generated_%s_marshal_json_%s.go", tt.typeName, method))
					if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
						t.Errorf("Expected generated file %s was not created", expectedFile)
					}
				}
			}
		})
	}
}

func Test_generateForStruct(t *testing.T) {
	// Create a test AST
	const src = `package testgen
type TestStruct struct {
	Field string ` + "`json:\"field\"`" + `
}
`
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	// Create a temporary directory
	tmpParent := filepath.Join(os.TempDir(), "pregogen_test")
	tmpdir := filepath.Join(tmpParent, "src", "testgen")
	if err := os.MkdirAll(tmpdir, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpParent)

	tests := []struct {
		name     string
		node     *ast.File
		typeName string
		genType  string
		wantErr  bool
	}{
		{
			name:     "Valid struct",
			node:     node,
			typeName: "TestStruct",
			genType:  "marshal",
			wantErr:  false,
		},
		{
			name:     "Non-existent type",
			node:     node,
			typeName: "NonExistentType",
			genType:  "marshal",
			wantErr:  false, // No error since it just returns
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := generateForStruct(tt.node, tt.typeName, tt.genType, tmpdir); (err != nil) != tt.wantErr {
				t.Errorf("generateForStruct() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && tt.typeName == "TestStruct" {
				// Verify files were generated in the correct directory
				for _, method := range GetMethods() {
					expectedFile := filepath.Join(tmpdir, fmt.Sprintf("generated_%s_marshal_json_%s.go", tt.typeName, method))
					if _, err := os.Stat(expectedFile); os.IsNotExist(err) {
						t.Errorf("Expected generated file %s was not created", expectedFile)
					}
				}
			}
		})
	}
}

func Test_generateJSON(t *testing.T) {
	// Create a test TypeSpec
	typeSpec := &ast.TypeSpec{
		Name: &ast.Ident{Name: "TestStruct"},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: "Field"}},
						Type:  &ast.Ident{Name: "string"},
						Tag:   &ast.BasicLit{Kind: token.STRING, Value: "`json:\"field\"`"},
					},
				},
			},
		},
	}

	// Create a temporary directory
	tmpParent := filepath.Join(os.TempDir(), "pregogen_test")
	tmpdir := filepath.Join(tmpParent, "src", "testgen")
	if err := os.MkdirAll(tmpdir, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpParent)

	tests := []struct {
		name        string
		typeSpec    *ast.TypeSpec
		packageName string
		gentype     string
		wantPanic   bool
	}{
		{
			name:        "Valid struct marshal",
			typeSpec:    typeSpec,
			packageName: "testgen",
			gentype:     "marshal",
			wantPanic:   false,
		},
		{
			name:        "Valid struct unmarshal",
			typeSpec:    typeSpec,
			packageName: "testgen",
			gentype:     "unmarshal",
			wantPanic:   false,
		},
		{
			name:        "Valid struct test",
			typeSpec:    typeSpec,
			packageName: "testgen",
			gentype:     "testMarshal",
			wantPanic:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("generateJSON() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()
			generateJSON(tt.typeSpec, tt.packageName, tt.gentype, tmpdir)

			if !tt.wantPanic {
				var expectedFiles []string
				if tt.gentype == "marshal" {
					// Check all method files
					for _, method := range GetMethods() {
						expectedFiles = append(expectedFiles,
							filepath.Join(tmpdir, fmt.Sprintf("generated_%s_marshal_json_%s.go", tt.typeSpec.Name.Name, method)))
					}
				} else if tt.gentype == "unmarshal" {
					expectedFiles = append(expectedFiles,
						filepath.Join(tmpdir, fmt.Sprintf("generated_%s_unmarshal_json.go", tt.typeSpec.Name.Name)))
				} else if tt.gentype == "testMarshal" {
					expectedFiles = append(expectedFiles,
						filepath.Join(tmpdir, fmt.Sprintf("generated_%s_json_test.go", tt.typeSpec.Name.Name)))
				}

				for _, file := range expectedFiles {
					if _, err := os.Stat(file); os.IsNotExist(err) {
						t.Errorf("Expected generated file %s was not created", file)
					}
				}
			}
		})
	}
}

func Test_processGentype(t *testing.T) {
	// Create a temporary directory
	tmpParent := filepath.Join(os.TempDir(), "pregogen_test")
	tmpdir := filepath.Join(tmpParent, "src", "testgen")
	if err := os.MkdirAll(tmpdir, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpParent)

	fields := []FieldInfo{
		{
			Name:         "Field",
			JSONName:     "field",
			Properties:   nil,
			FieldType:    "string",
			ReceiverName: "t_receiver",
		},
	}

	tests := []struct {
		name        string
		gentype     string
		packageName string
		wantPanic   bool
	}{
		{
			name:        "Valid bytesBuffer",
			gentype:     "bytesBuffer",
			packageName: "testgen",
			wantPanic:   false,
		},
		{
			name:        "Valid stringsBuilder",
			gentype:     "stringsBuilder",
			packageName: "testgen",
			wantPanic:   false,
		},
		{
			name:        "Valid append",
			gentype:     "append",
			packageName: "testgen",
			wantPanic:   false,
		},
		{
			name:        "Valid plus",
			gentype:     "plus",
			packageName: "testgen",
			wantPanic:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("processGentype() panic = %v, wantPanic %v", r, tt.wantPanic)
				}
			}()
			filename := filepath.Join(tmpdir, "test_"+tt.gentype+".go")
			processGentype(tt.gentype, tt.packageName, "t_receiver", "TestType", filename, fields, nil)

			if !tt.wantPanic {
				if _, err := os.Stat(filename); os.IsNotExist(err) {
					t.Errorf("Expected generated file %s was not created", filename)
				}
			}
		})
	}
}
