// Package main implements a code generator for adding custom JSON encoding and decoding methods
// to Go structs without using the standard json library. This generator uses go generate to
// parse Go source files, identify structs, and generate the necessary methods based on a template.
package pregogen

import (
	"go/ast"
	"reflect"
	"testing"
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
