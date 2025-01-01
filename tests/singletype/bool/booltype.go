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
	string
}{
	{
		BoolType{
			BoolField: true,
		}, `{"boolfield":true}`},
	{
		BoolType{
			BoolField: false,
		}, `{"boolfield":false}`},
}
var BoolType_example = BoolType_examples[0].BoolType

type BoolType3 struct {
	BoolField1 bool `json:"boolfield1"`
	BoolField2 bool `json:"boolfield2"`
	BoolField3 bool `json:"boolfield3"`
}

// representative example(s) of data stored in target application
var BoolType3_examples = []struct {
	BoolType3
	string
}{
	{
		BoolType3{
			BoolField1: true,
			BoolField2: true,
			BoolField3: false,
		}, `{"boolfield1":true,"boolfield2":true,"boolfield3":false}`},
	{
		BoolType3{
			BoolField1: false,
			BoolField2: false,
			BoolField3: true,
		}, `{"boolfield1":false,"boolfield2":false,"boolfield3":true}`},
}
var BoolType3_example = BoolType3_examples[0].BoolType3
