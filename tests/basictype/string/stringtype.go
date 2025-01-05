//go:generate pregogen -type=StringType -file=$GOFILE -gen=test
//go:generate pregogen -type=StringType -file=$GOFILE -gen=append
//go:generate pregogen -type=StringType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=StringType -file=$GOFILE -gen=plus

package stringtype

type StringType struct {
	StringField string `json:"stringfield"`
}

// representative example of data stored in target application
var StringType_examples = []struct {
	StringType
	want []byte
}{
	{
		StringType{
			StringField: "hello",
		}, nil},
}

//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=plus
type StringType3 struct {
	StringField1 string `json:"stringfield1"`
	StringField2 string `json:"stringfield2"`
	StringField3 string `json:"stringfield3"`
}

// representative example(s) of data stored in target application
var StringType3_examples = []struct {
	StringType3
	want []byte
}{
	{
		StringType3{
			StringField1: "true",
			StringField2: "true",
			StringField3: "false",
		}, nil},
}
