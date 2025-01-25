//go:generate pregogen -type=BoolType -file=$GOFILE -gen=test
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=append
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=BoolType -file=$GOFILE -gen=plus

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

//go:generate pregogen -type=BoolType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=BoolType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=BoolType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=BoolType3 -file=$GOFILE -gen=plus

// representative example(s) of data stored in target application
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

//go:generate pregogen -type=BoolArrayType -file=$GOFILE -gen=test
//go:generate pregogen -type=BoolArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=BoolArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=BoolArrayType -file=$GOFILE -gen=plus
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
}

type BoolArrayType3 struct {
	BoolArrayField1 []bool `json:"boolarray1field"`
	BoolArrayField2 []bool `json:"boolarray2field"`
	BoolArrayField3 []bool `json:"boolarray3field"`
}

//go:generate pregogen -type=BoolArrayType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=BoolArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=BoolArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=BoolArrayType3 -file=$GOFILE -gen=plus

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
}
