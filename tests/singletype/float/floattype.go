//go:generate pregogen -type=FloatType -file=floattype.go -gen=test
//go:generate pregogen -type=FloatType -file=floattype.go -gen=append
//go:generate pregogen -type=FloatType -file=floattype.go -gen=bytesBuffer
//go:generate pregogen -type=FloatType -file=floattype.go -gen=plus

//go:generate pregogen -type=FloatType3 -file=floattype.go -gen=test
//go:generate pregogen -type=FloatType3 -file=floattype.go -gen=append
//go:generate pregogen -type=FloatType3 -file=floattype.go -gen=bytesBuffer
//go:generate pregogen -type=FloatType3 -file=floattype.go -gen=plus
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
		}, []byte(`{"floatfield":12.45}`)},
}

type FloatType3 struct {
	FloatField1 float64 `json:"Floatfield1"`
	FloatField2 float64 `json:"Floatfield2"`
	FloatField3 float64 `json:"Floatfield3"`
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
		}, []byte(`{"floatfield1":123.45,"floatfield2":678.9,"floatfield3":1.0}`)},
}
