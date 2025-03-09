// Package main implements a code generator for adding custom JSON encoding and decoding methods
// to Go structs without using the standard json library. This generator uses go generate to
// parse Go source files, identify structs, and generate the necessary methods based on a template.
package pregogen

import (
	"go/ast"
	"reflect"
	"testing"
	"time"
)

func Test_addIfNotPresent(t *testing.T) {
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Test new", args{[]string{"a", "b", "c"}, "d"}, false},
		{"Test a", args{[]string{"a", "b", "c"}, "a"}, true},
		{"Test b", args{[]string{"a", "b", "c"}, "b"}, true},
		{"Test c", args{[]string{"a", "b", "c"}, "c"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsItem(tt.args.slice, tt.args.item)
			if got != tt.want {
				t.Errorf("addIfNotPresent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseJSONTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name           string
		args           args
		wantName       string
		wantProperties map[string]bool
	}{
		// TODO: Add test cases.
		{"Test 0", args{""}, "", nil},
		{"Test 00", args{"toto:\"stringfield\""}, "", map[string]bool{}},
		{"Test 000", args{"toto:\"stringfield,omitempty\""}, "", map[string]bool{}},
		{"Test 1", args{"json:\"stringfield\""}, "stringfield", map[string]bool{}},
		{"Test 2", args{"json:\"stringfield,omitempty\""}, "stringfield", map[string]bool{"omitempty": true}},
		{"Test 3", args{"json:\"stringfield,omitempty,required\""}, "stringfield", map[string]bool{"omitempty": true, "required": true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, gotProperties := parseJSONTag(tt.args.tag)
			if gotName != tt.wantName {
				t.Errorf("parseJSONTag() got = %v, want %v", gotName, tt.wantName)
			}
			if !reflect.DeepEqual(gotProperties, tt.wantProperties) {
				t.Errorf("parseJSONTag() got1 = %v, want %v", gotProperties, tt.wantProperties)
			}
		})
	}
}

func Test_getTypeName(t *testing.T) {
	type args struct {
		expr ast.Expr
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test 0", args{&ast.Ident{Name: "string"}}, "string"},
		{"Test 1", args{&ast.ArrayType{Elt: &ast.Ident{Name: "string"}}}, "[]string"},
		{"Test 2", args{&ast.StarExpr{X: &ast.Ident{Name: "string"}}}, "*string"},
		{"Test 3", args{&ast.SelectorExpr{X: &ast.Ident{Name: "mypac"}, Sel: &ast.Ident{Name: "string"}}}, "mypac.string"},
		// test default
		{"Test 4", args{&ast.BasicLit{Value: "string"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTypeName(tt.args.expr); got != tt.want {
				t.Errorf("getTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processFieldType(t *testing.T) {
	type args struct {
		fieldType string
		fieldName string
		includes  []string
		want      []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"Test string", args{"string", "StringField", []string{}, []string{}}},
		{"Test date", args{"time.Time", "DateField", []string{}, []string{}}},
		{"Test int", args{"int", "IntField", []string{}, []string{"strconv"}}},
		{"Test second int", args{"int", "IntField", []string{"strconv", "encoding/json"}, []string{"strconv", "encoding/json"}}},
		{"Test second int bis", args{"int", "IntField", []string{"strconv"}, []string{"strconv"}}},
		{"Test int8", args{"int8", "Int8Field", []string{}, []string{}}},
		{"Test byte", args{"byte", "ByteField", []string{}, []string{}}},
		{"Test bool", args{"bool", "BoolField", []string{}, []string{}}},
		{"Test uint16", args{"uint16", "Uint16Field", []string{}, []string{"strconv"}}},
		{"Test float64", args{"float64", "FloatField", []string{}, []string{"strconv"}}},
		{"Test *string", args{"*string", "PointerStringField", []string{}, []string{}}},
		{"Test []string", args{"[]string", "StringArrayField", []string{}, []string{}}},
		{"Test []*string", args{"[]*string", "PointerStringArrayField", []string{}, []string{}}},
		{"Test other", args{"rune", "RuneField", []string{}, []string{"encoding/json"}}},
		{"Test second other", args{"rune", "RuneField", []string{"encoding/json"}, []string{"encoding/json"}}},
		{"Test second other bis", args{"rune", "RuneField", []string{"strconv", "encoding/json"}, []string{"strconv", "encoding/json"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := tt.args.includes
			processFieldType(tt.args.fieldType, tt.args.fieldName, &before)
			if !reflect.DeepEqual(before, tt.args.want) {
				t.Errorf("processFieldType() = %v, want %v", before, tt.args.want)
			}
		})
	}
}

func Test_dict(t *testing.T) {
	tests := []struct {
		name    string
		values  []interface{}
		want    map[string]interface{}
		wantErr bool
	}{
		{"Valid input", []interface{}{"key1", "value1", "key2", "value2"}, map[string]interface{}{"key1": "value1", "key2": "value2"}, false},
		{"Invalid input", []interface{}{"key1", "value1", "key2"}, nil, true},
		{"Non-string key", []interface{}{1, "value1"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dict(tt.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("dict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Sub(t *testing.T) {
	if got := Sub(5, 3); got != 2 {
		t.Errorf("Sub() = %v, want %v", got, 2)
	}
}

func Test_GetArrayType(t *testing.T) {
	if got := GetArrayType("[]string"); got != "string" {
		t.Errorf("GetArrayType() = %v, want %v", got, "string")
	}
}

func Test_GetPointerType(t *testing.T) {
	if got := GetPointerType("*string"); got != "string" {
		t.Errorf("GetPointerType() = %v, want %v", got, "string")
	}
}

func Test_IsArray(t *testing.T) {
	if got := IsArray("[]string"); got != true {
		t.Errorf("IsArray() = %v, want %v", got, true)
	}
}

func Test_Seq(t *testing.T) {
	if got := Seq(1, 3); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Seq() = %v, want %v", got, []int{1, 2, 3})
	}
}

func Test_GetMethods(t *testing.T) {
	if got := GetMethods(); !reflect.DeepEqual(got, []string{"bytesBuffer", "stringsBuilder", "append", "plus"}) {
		t.Errorf("GetMethods() = %v, want %v", got, []string{"bytesBuffer", "stringsBuilder", "append", "plus"})
	}
}

func Test_Method(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		action   string
		addition string
		want     string
		wantErr  bool
	}{
		{"bytesBuffer", "bytesBuffer", "start", "addition", "result.WriteString(addition", false},
		{"stringsBuilder", "stringsBuilder", "start", "addition", "result.WriteString(addition", false},
		{"append", "append", "start", "addition", "result = append(result, addition", false},
		{"plus", "plus", "start", "addition", "result += addition", false},
		{"unknown", "unknown", "start", "addition", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Method(tt.method, tt.action, tt.addition)
			if (err != nil) != tt.wantErr {
				t.Errorf("Method() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Method() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MethodAppend(t *testing.T) {
	tests := []struct {
		name     string
		action   string
		addition string
		want     string
		wantErr  bool
	}{
		{"declareVar", "declareVar", "", "\n\tvar result []byte\n\t", false},
		{"start", "start", "addition", "result = append(result, addition", false},
		{"stop", "stop", "addition", "addition...)\n\t", false},
		{"unknown", "unknown", "addition", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MethodAppend(tt.action, tt.addition)
			if (err != nil) != tt.wantErr {
				t.Errorf("MethodAppend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MethodAppend() = <<%v>>, want <<%v>>", got, tt.want)
			}
		})
	}
}

func Test_MethodBytesBuffer(t *testing.T) {
	tests := []struct {
		name     string
		action   string
		addition string
		want     string
		wantErr  bool
	}{
		{"declareVar", "declareVar", "", "\n\tvar result bytes.Buffer\n\t", false},
		{"start", "start", "addition", "result.WriteString(addition", false},
		{"stop", "stop", "addition", "addition)\n\t", false},
		{"unknown", "unknown", "addition", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MethodBytesBuffer(tt.action, tt.addition)
			if (err != nil) != tt.wantErr {
				t.Errorf("MethodBytesBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MethodBytesBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MethodStringsBuilder(t *testing.T) {
	tests := []struct {
		name     string
		action   string
		addition string
		want     string
		wantErr  bool
	}{
		{"declareVar", "declareVar", "", "\n\tvar result strings.Builder\n\t", false},
		{"start", "start", "addition", "result.WriteString(addition", false},
		{"stop", "stop", "addition", "addition)\n\t", false},
		{"unknown", "unknown", "addition", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MethodStringsBuilder(tt.action, tt.addition)
			if (err != nil) != tt.wantErr {
				t.Errorf("MethodStringsBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MethodStringsBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MethodPlus(t *testing.T) {
	tests := []struct {
		name     string
		action   string
		addition string
		want     string
		wantErr  bool
	}{
		{"declareVar", "declareVar", "", "\n\tvar result string\n\t", false},
		{"start", "start", "addition", "result += addition", false},
		{"stop", "stop", "addition", "addition\n\t", false},
		{"unknown", "unknown", "addition", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MethodPlus(tt.action, tt.addition)
			if (err != nil) != tt.wantErr {
				t.Errorf("MethodPlus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MethodPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MethodUnique(t *testing.T) {
	tests := []struct {
		name     string
		action   string
		addition string
		want     string
		wantErr  bool
	}{
		{"start", "start", "addition", "return ([]byte(", false},
		{"stop", "stop", "addition", ")), nil", false},
		{"unknown", "unknown", "addition", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MethodUnique(tt.action, tt.addition)
			if (err != nil) != tt.wantErr {
				t.Errorf("MethodUnique() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MethodUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Return(t *testing.T) {
	// Use a channel to coordinate the test
	done := make(chan bool)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				// We expect a panic with "exit"
				if r != "exit" {
					t.Errorf("Return() unexpected panic value: %v", r)
				}
				done <- true
			}
		}()
		Return()
		t.Error("Return() should not reach this point")
	}()

	// Wait for the goroutine to complete
	select {
	case <-done:
		// Test passed
	case <-time.After(time.Second):
		t.Error("Test_Return timed out")
	}
}

func Test_templateFuncs(t *testing.T) {
	if got := templateFuncs(); got == nil {
		t.Errorf("templateFuncs() = %v, want non-nil", got)
	}
}
