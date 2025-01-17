//go:generate pregogen -type=FloatType -file=$GOFILE -gen=test
//go:generate pregogen -type=FloatType -file=$GOFILE -gen=append
//go:generate pregogen -type=FloatType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=FloatType -file=$GOFILE -gen=plus

package floattype

type FloatType struct {
	FloatField float64 `json:"floatfield"`
}

// representative example of data stored in target application
var FloatType_examples = []struct {
	FloatType
	want []byte
}{
	{
		FloatType{
			FloatField: 12.45,
		}, nil},
}

//go:generate pregogen -type=FloatType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=FloatType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=FloatType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=FloatType3 -file=$GOFILE -gen=plus
type FloatType3 struct {
	FloatField1 float64 `json:"floatfield1"`
	FloatField2 float64 `json:"floatfield2"`
	FloatField3 float64 `json:"floatfield3"`
}

// representative example(s) of data stored in target application
var FloatType3_examples = []struct {
	FloatType3
	want []byte
}{
	{
		FloatType3{
			FloatField1: 123.45,
			FloatField2: 678.9,
			FloatField3: 1.0,
		}, nil},
}
