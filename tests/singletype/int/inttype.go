//go:generate pregogen -type=IntType -file=inttype.go -gen=test
//go:generate pregogen -type=IntType -file=inttype.go -gen=append
//go:generate pregogen -type=IntType -file=inttype.go -gen=bytesBuffer
//go:generate pregogen -type=IntType -file=inttype.go -gen=plus

//go:generate pregogen -type=IntType3 -file=inttype.go -gen=test
//go:generate pregogen -type=IntType3 -file=inttype.go -gen=append
//go:generate pregogen -type=IntType3 -file=inttype.go -gen=bytesBuffer
//go:generate pregogen -type=IntType3 -file=inttype.go -gen=plus
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
		}, []byte(`{"intfield":12345}`)},
}

type IntType3 struct {
	IntField1 int `json:"Intfield1"`
	IntField2 int `json:"Intfield2"`
	IntField3 int `json:"Intfield3"`
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
		}, []byte(`{"intfield1":12345,"intfield2":67890,"intfield3":10}`)},
}
