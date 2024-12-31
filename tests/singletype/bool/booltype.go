//go:generate pregogen -type=BoolType -file=booltype.go -gen=append
package booltype

type BoolType struct {
	BoolField bool `json:"boolfield"`
}

// representative example of data stored in target application
var BoolType_example = BoolType{
	BoolField: true,
}
