//go:generate pregogen -type=FloatType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=FloatType -file=$GOFILE -gen=marshal
//go:generate pregogen -type=FloatType -file=$GOFILE -gen=unmarshal
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

//go:generate pregogen -type=FloatType3 -file=$GOFILE -gen=testAll
//go:generate pregogen -type=FloatType3 -file=$GOFILE -gen=marshal
//go:generate pregogen -type=FloatType3 -file=$GOFILE -gen=unmarshal
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

//go:generate pregogen -type=Float32ArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=Float32ArrayType -file=$GOFILE -gen=marshal
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

//go:generate pregogen -type=Float32ArrayType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=Float32ArrayType3 -file=$GOFILE -gen=marshal

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

//go:generate pregogen -type=PointerFloat32Type -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerFloat32Type -file=$GOFILE -gen=marshal

type PointerFloat32Type struct {
	PointerFloat32Field *float32 `json:"pointerfloat32field"`
}

var int1 float32 = 12345.12

// representative example of data stored in target application
var PointerFloat32Type_examples = []struct {
	PointerFloat32Type
	want []byte
}{
	{
		PointerFloat32Type{
			PointerFloat32Field: &int1,
		}, nil},
	{
		PointerFloat32Type{
			PointerFloat32Field: nil,
		}, nil},
}

//go:generate pregogen -type=PointerFloat32Type3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerFloat32Type3 -file=$GOFILE -gen=marshal
type PointerFloat32Type3 struct {
	PointerFloat32Field1 *float32 `json:"pointerfloat32field1"`
	PointerFloat32Field2 *float32 `json:"pointerfloat32field2"`
	PointerFloat32Field3 *float64 `json:"pointerfloat32field3"`
}

var int2 float32 = 12345.14
var int3 float64 = 67890.23

// representative example(s) of data stored in target application
var PointerFloat32Type3_examples = []struct {
	PointerFloat32Type3
	want []byte
}{
	{
		PointerFloat32Type3{
			PointerFloat32Field1: &int1,
			PointerFloat32Field2: nil,
			PointerFloat32Field3: &int3,
		}, nil},
	{
		PointerFloat32Type3{
			PointerFloat32Field1: nil,
			PointerFloat32Field2: nil,
			PointerFloat32Field3: nil,
		}, nil},
	{
		PointerFloat32Type3{
			PointerFloat32Field1: nil,
			PointerFloat32Field2: &int2,
			PointerFloat32Field3: nil,
		}, nil},
}

//go:generate pregogen -type=PointerFloatArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerFloatArrayType -file=$GOFILE -gen=marshal

type PointerFloatArrayType struct {
	PointerFloatField []*float32 `json:"pointerfloatfield"`
}

var intarray1 []*float32 = []*float32{&int1, nil, &int2}
var intarray2 []*float32 = []*float32{nil, nil, nil}

// representative example of data stored in target application
var PointerFloatArrayType_examples = []struct {
	PointerFloatArrayType
	want []byte
}{
	{
		PointerFloatArrayType{
			PointerFloatField: intarray1,
		}, nil},
	{
		PointerFloatArrayType{
			PointerFloatField: intarray2,
		}, nil},
}

//go:generate pregogen -type=PointerFloatArrayType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerFloatArrayType3 -file=$GOFILE -gen=marshal
type PointerFloatArrayType3 struct {
	PointerFloatArrayField1 []*float32 `json:"pointerfloatfield1"`
	PointerFloatArrayField2 []*float32 `json:"pointerfloatfield2"`
	PointerFloatArrayField3 []*float64 `json:"pointerfloatfield3"`
}

var intarray3 []*float64 = []*float64{nil, &int3, nil}

// representative example(s) of data stored in target application
var PointerFloatArrayType3_examples = []struct {
	PointerFloatArrayType3
	want []byte
}{
	{
		PointerFloatArrayType3{
			PointerFloatArrayField1: intarray1,
			PointerFloatArrayField2: nil,
			PointerFloatArrayField3: intarray3,
		}, nil},
	{
		PointerFloatArrayType3{
			PointerFloatArrayField1: nil,
			PointerFloatArrayField2: nil,
			PointerFloatArrayField3: nil,
		}, nil},
	{
		PointerFloatArrayType3{
			PointerFloatArrayField1: nil,
			PointerFloatArrayField2: intarray2,
			PointerFloatArrayField3: nil,
		}, nil},
}
