//go:generate pregogen -type=BoolType -file=booltype.go -gen=test
//go:generate pregogen -type=BoolType -file=booltype.go -gen=append
//go:generate pregogen -type=BoolType -file=booltype.go -gen=bytesBuffer
//go:generate pregogen -type=BoolType -file=booltype.go -gen=plus

//go:generate pregogen -type=BoolType3 -file=booltype.go -gen=test
//go:generate pregogen -type=BoolType3 -file=booltype.go -gen=append
//go:generate pregogen -type=BoolType3 -file=booltype.go -gen=bytesBuffer
//go:generate pregogen -type=BoolType3 -file=booltype.go -gen=plus
package booltype

type BoolType struct {
	BoolField bool `json:"boolfield"`
}

// representative example of data stored in target application
var BoolType_examples = []struct {
	BoolType
	want []byte
}{
	{
		BoolType{
			BoolField: true,
		}, nil},
	{
		BoolType{
			BoolField: false,
		}, nil},
}

type BoolType3 struct {
	BoolField1 bool `json:"boolfield1"`
	BoolField2 bool `json:"boolfield2"`
	BoolField3 bool `json:"boolfield3"`
}

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
