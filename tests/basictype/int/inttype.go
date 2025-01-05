//go:generate pregogen -type=IntType -file=$GOFILE -gen=test
//go:generate pregogen -type=IntType -file=$GOFILE -gen=append
//go:generate pregogen -type=IntType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=IntType -file=$GOFILE -gen=plus

package inttype

type IntType struct {
	IntField int `json:"intfield"`
}

// representative example of data stored in target application
var IntType_examples = []struct {
	IntType
	want []byte
}{
	{
		IntType{
			IntField: 12345,
		}, nil},
}

//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=plus
type IntType3 struct {
	IntField1 int `json:"intfield1"`
	IntField2 int `json:"intfield2"`
	IntField3 int `json:"intfield3"`
}

// representative example(s) of data stored in target application
var IntType3_examples = []struct {
	IntType3
	want []byte
}{
	{
		IntType3{
			IntField1: 12345,
			IntField2: 67890,
			IntField3: 10,
		}, nil},
}
