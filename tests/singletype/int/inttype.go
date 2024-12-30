//go:generate pregogen -type=IntType -file=inttype.go -gen=plus
package inttype

type IntType struct {
	IntField int `json:"intfield"`
}

// representative example of data stored in target application
var IntType_example = IntType{
	IntField: 12345,
}
