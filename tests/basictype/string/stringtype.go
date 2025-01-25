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

//go:generate pregogen -type=StringArrayType -file=$GOFILE -gen=test
//go:generate pregogen -type=StringArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=StringArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=StringArrayType -file=$GOFILE -gen=plus
type StringArrayType struct {
	StringArrayField []string `json:"stringarrayfield"`
}

// representative example of data stored in target application
var StringArrayType_examples = []struct {
	StringArrayType StringArrayType
	want            []byte
}{
	{
		StringArrayType{
			StringArrayField: []string{"1", "2", "3"},
		}, nil},
}

type StringArrayType3 struct {
	StringArrayField1 []string `json:"stringarray1field"`
	StringArrayField2 []string `json:"stringarray2field"`
	StringArrayField3 []string `json:"stringarray3field"`
}

//go:generate pregogen -type=StringArrayType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=StringArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=StringArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=StringArrayType3 -file=$GOFILE -gen=plus

// representative example(s) of data stored in target application
var StringArrayType3_examples = []struct {
	StringArrayType3
	want []byte
}{
	{
		StringArrayType3{
			StringArrayField1: []string{"1", "2", "3"},
			StringArrayField2: []string{"4", "5", "6"},
			StringArrayField3: []string{"7", "8", "9"},
		}, nil},
}
