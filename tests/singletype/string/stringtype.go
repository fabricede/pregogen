//go:generate pregogen -type=StringType -file=stringtype.go -gen=test
//go:generate pregogen -type=StringType -file=stringtype.go -gen=append
//go:generate pregogen -type=StringType -file=stringtype.go -gen=bytesBuffer
//go:generate pregogen -type=StringType -file=stringtype.go -gen=plus

//go:generate pregogen -type=StringType3 -file=stringtype.go -gen=test
//go:generate pregogen -type=StringType3 -file=stringtype.go -gen=append
//go:generate pregogen -type=StringType3 -file=stringtype.go -gen=bytesBuffer
//go:generate pregogen -type=StringType3 -file=stringtype.go -gen=plus
package stringtype

type StringType struct {
	StringField string `json:"stringfield"`
}

// representative example of data stored in target application
var StringType_examples = []struct {
	StringType
	string
}{
	{
		StringType{
			StringField: "hello",
		}, `{"stringfield":"hello"}`},
}
var StringType_example = StringType_examples[0].StringType

type StringType3 struct {
	StringField1 string `json:"Stringfield1"`
	StringField2 string `json:"Stringfield2"`
	StringField3 string `json:"Stringfield3"`
}

// representative example(s) of data stored in target application
var StringType3_examples = []struct {
	StringType3
	string
}{
	{
		StringType3{
			StringField1: "true",
			StringField2: "true",
			StringField3: "false",
		}, `{"stringfield1":"true","stringfield2":"true","stringfield3":"false"}`},
}
var StringType3_example = StringType3_examples[0].StringType3
