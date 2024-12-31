//go:generate pregogen -type=StringType -file=stringtype.go -gen=append
package stringtype

type StringType struct {
	StringField string `json:"stringfield"`
}

// representative example of data stored in target application
var StringType_example = StringType{
	StringField: "hello",
}
