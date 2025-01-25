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

//go:generate pregogen -type=Float32ArrayType -file=$GOFILE -gen=test
//go:generate pregogen -type=Float32ArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=Float32ArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=Float32ArrayType -file=$GOFILE -gen=plus
type Float32ArrayType struct {
	Float32ArrayField []float32 `json:"float32arrayfield"`
}

// representative example of data stored in target application
var Float32ArrayType_examples = []struct {
	Float32ArrayType Float32ArrayType
	want             []byte
}{
	{
		Float32ArrayType{
			Float32ArrayField: []float32{1, 2, 3},
		}, nil},
}

type Float32ArrayType3 struct {
	Float32ArrayField1 []float32 `json:"float32array1field"`
	Float32ArrayField2 []float32 `json:"float32array2field"`
	Float32ArrayField3 []float32 `json:"float32array3field"`
}

//go:generate pregogen -type=Float32ArrayType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=Float32ArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=Float32ArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=Float32ArrayType3 -file=$GOFILE -gen=plus

// representative example(s) of data stored in target application
var Float32ArrayType3_examples = []struct {
	Float32ArrayType3
	want []byte
}{
	{
		Float32ArrayType3{
			Float32ArrayField1: []float32{1, 2, 3},
			Float32ArrayField2: []float32{4, 5, 6},
			Float32ArrayField3: []float32{7, 8, 9},
		}, nil},
}
