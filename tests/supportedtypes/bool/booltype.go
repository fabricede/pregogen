// in the case of uniquebool we can't use marshal due to return in template
// details all methods instead
//
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=append
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=plus
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=stringsBuilder
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=unmarshal
package booltype

type BoolType struct {
	BoolField bool `json:"boolfield"`
}

// representative example of data stored in target application
var BoolType_examples = []struct {
	BoolType BoolType
	want     []byte
}{
	{BoolType{BoolField: true}, nil},
	{BoolType{BoolField: false}, nil},
}

type BoolType3 struct {
	BoolField1 bool `json:"boolfield1"`
	BoolField2 bool `json:"boolfield2"`
	BoolField3 bool `json:"boolfield3"`
}

// representative example(s) of data stored in target application
//
//go:generate pregogen -type=BoolType3 -file=$GOFILE -gen=testAll
//go:generate pregogen -type=BoolType3 -file=$GOFILE -gen=marshal
//go:generate pregogen -type=BoolType3 -file=$GOFILE -gen=unmarshal
var BoolType3_examples = []struct {
	BoolType3
	want []byte
}{
	{
		BoolType3{
			BoolField1: true,
			BoolField2: true,
			BoolField3: false,
		}, nil},
	{
		BoolType3{
			BoolField1: false,
			BoolField2: false,
			BoolField3: true,
		}, nil},
}

//go:generate pregogen -type=BoolArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=BoolArrayType -file=$GOFILE -gen=marshal
type BoolArrayType struct {
	BoolArrayField []bool `json:"boolarrayfield"`
}

// representative example of data stored in target application
var BoolArrayType_examples = []struct {
	BoolArrayType BoolArrayType
	want          []byte
}{
	{
		BoolArrayType{
			BoolArrayField: []bool{true, true, false},
		}, nil},
	{
		BoolArrayType{
			BoolArrayField: []bool{false, false, true},
		}, nil},
	{
		BoolArrayType{
			BoolArrayField: []bool{true},
		}, nil},
	{
		BoolArrayType{
			BoolArrayField: []bool{false},
		}, nil},
}

type BoolArrayType3 struct {
	BoolArrayField1 []bool `json:"boolarray1field"`
	BoolArrayField2 []bool `json:"boolarray2field"`
	BoolArrayField3 []bool `json:"boolarray3field"`
}

//go:generate pregogen -type=BoolArrayType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=BoolArrayType3 -file=$GOFILE -gen=marshal

// representative example(s) of data stored in target application
var BoolArrayType3_examples = []struct {
	BoolArrayType3
	want []byte
}{
	{
		BoolArrayType3{
			BoolArrayField1: []bool{true, true, false},
			BoolArrayField2: []bool{false, true, true},
			BoolArrayField3: []bool{true, false, false},
		}, nil},
	{
		BoolArrayType3{
			BoolArrayField1: []bool{false, false, true},
			BoolArrayField2: []bool{true, false, false},
			BoolArrayField3: []bool{false, true, true},
		}, nil},
	{
		BoolArrayType3{
			BoolArrayField1: []bool{true},
			BoolArrayField2: []bool{false},
			BoolArrayField3: []bool{true},
		}, nil},
	{
		BoolArrayType3{
			BoolArrayField1: []bool{false},
			BoolArrayField2: []bool{true},
			BoolArrayField3: []bool{false},
		}, nil},
}

//go:generate pregogen -type=PointerBoolType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerBoolType -file=$GOFILE -gen=marshal

type PointerBoolType struct {
	PointerBoolField *bool `json:"pointerboolfield"`
}

var bool1 = true
var bool2 = false

// representative example of data stored in target application
var PointerBoolType_examples = []struct {
	PointerBoolType
	want []byte
}{
	{
		PointerBoolType{
			PointerBoolField: &bool1,
		}, nil},
	{
		PointerBoolType{
			PointerBoolField: &bool2,
		}, nil},
	{
		PointerBoolType{
			PointerBoolField: nil,
		}, nil},
}

//go:generate pregogen -type=PointerBoolType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerBoolType3 -file=$GOFILE -gen=marshal
type PointerBoolType3 struct {
	PointerBoolField1 *bool `json:"pointerboolfield1"`
	PointerBoolField2 *bool `json:"pointerboolfield2"`
	PointerBoolField3 *bool `json:"pointerboolfield3"`
}

var bool3 = true

// representative example(s) of data stored in target application
var PointerBoolType3_examples = []struct {
	PointerBoolType3
	want []byte
}{
	{
		PointerBoolType3{
			PointerBoolField1: &bool1,
			PointerBoolField2: nil,
			PointerBoolField3: &bool3,
		}, nil},
	{
		PointerBoolType3{
			PointerBoolField1: nil,
			PointerBoolField2: nil,
			PointerBoolField3: nil,
		}, nil},
	{
		PointerBoolType3{
			PointerBoolField1: nil,
			PointerBoolField2: &bool2,
			PointerBoolField3: nil,
		}, nil},
}

//go:generate pregogen -type=PointerBoolArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerBoolArrayType -file=$GOFILE -gen=marshal

type PointerBoolArrayType struct {
	PointerBoolField []*bool `json:"pointerboolfield"`
}

var bytearray1 []*bool = []*bool{&bool1, nil, &bool3}
var bytearray2 []*bool = []*bool{nil, nil, nil}

// representative example of data stored in target application
var PointerBoolArrayType_examples = []struct {
	PointerBoolArrayType
	want []byte
}{
	{
		PointerBoolArrayType{
			PointerBoolField: bytearray1,
		}, nil},
	{
		PointerBoolArrayType{
			PointerBoolField: bytearray2,
		}, nil},
}

//go:generate pregogen -type=PointerBoolArrayType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerBoolArrayType3 -file=$GOFILE -gen=marshal
type PointerBoolArrayType3 struct {
	PointerBoolArrayField1 []*bool `json:"pointerboolfield1"`
	PointerBoolArrayField2 []*bool `json:"pointerboolfield2"`
	PointerBoolArrayField3 []*bool `json:"pointerboolfield3"`
}

var bytearray3 []*bool = []*bool{nil, &bool3, nil}

// representative example(s) of data stored in target application
var PointerBoolArrayType3_examples = []struct {
	PointerBoolArrayType3
	want []byte
}{
	{
		PointerBoolArrayType3{
			PointerBoolArrayField1: bytearray1,
			PointerBoolArrayField2: nil,
			PointerBoolArrayField3: bytearray3,
		}, nil},
	{
		PointerBoolArrayType3{
			PointerBoolArrayField1: nil,
			PointerBoolArrayField2: nil,
			PointerBoolArrayField3: nil,
		}, nil},
	{
		PointerBoolArrayType3{
			PointerBoolArrayField1: nil,
			PointerBoolArrayField2: bytearray2,
			PointerBoolArrayField3: nil,
		}, nil},
}
