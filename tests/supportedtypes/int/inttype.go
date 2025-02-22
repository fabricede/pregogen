//go:generate pregogen -type=IntType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=IntType -file=$GOFILE -gen=append
//go:generate pregogen -type=IntType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=IntType -file=$GOFILE -gen=plus

package inttype

type IntType struct {
	IntField int `json:"intfield"`
}

// representative example of data stored in target application
var IntType_examples = []struct {
	IntType
	want []byte
}{
	{
		IntType{
			IntField: 12345,
		}, nil},
}

//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=plus
type IntType3 struct {
	IntField1 int `json:"intfield1"`
	IntField2 int `json:"intfield2"`
	IntField3 int `json:"intfield3"`
}

// representative example(s) of data stored in target application
var IntType3_examples = []struct {
	IntType3
	want []byte
}{
	{
		IntType3{
			IntField1: 12345,
			IntField2: 67890,
			IntField3: 10,
		}, nil},
}

//go:generate pregogen -type=IntArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=IntArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=IntArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=IntArrayType -file=$GOFILE -gen=plus
type IntArrayType struct {
	IntArrayField []int `json:"intarrayfield"`
}

// representative example of data stored in target application
var IntArrayType_examples = []struct {
	IntArrayType IntArrayType
	want         []byte
}{
	{
		IntArrayType{
			IntArrayField: []int{1, 2, 3},
		}, nil},
}

type IntArrayType3 struct {
	IntArrayField1 []int `json:"intarray1field"`
	IntArrayField2 []int `json:"intarray2field"`
	IntArrayField3 []int `json:"intarray3field"`
}

//go:generate pregogen -type=IntArrayType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=IntArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=IntArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=IntArrayType3 -file=$GOFILE -gen=plus

// representative example(s) of data stored in target application
var IntArrayType3_examples = []struct {
	IntArrayType3
	want []byte
}{
	{
		IntArrayType3{
			IntArrayField1: []int{1, 2, 3},
			IntArrayField2: []int{4, 5, 6},
			IntArrayField3: []int{7, 8, 9},
		}, nil},
}

//go:generate pregogen -type=PointerIntType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerIntType -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerIntType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerIntType -file=$GOFILE -gen=plus

type PointerIntType struct {
	PointerIntField *int `json:"pointerintfield"`
}

var int1 = 12345

// representative example of data stored in target application
var PointerIntType_examples = []struct {
	PointerIntType
	want []byte
}{
	{
		PointerIntType{
			PointerIntField: &int1,
		}, nil},
	{
		PointerIntType{
			PointerIntField: nil,
		}, nil},
}

//go:generate pregogen -type=PointerIntType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerIntType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerIntType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerIntType3 -file=$GOFILE -gen=plus
type PointerIntType3 struct {
	PointerIntField1 *int `json:"pointerintfield1"`
	PointerIntField2 *int `json:"pointerintfield2"`
	PointerIntField3 *int `json:"pointerintfield3"`
}

var int2 = 12345
var int3 = 67890

// representative example(s) of data stored in target application
var PointerIntType3_examples = []struct {
	PointerIntType3
	want []byte
}{
	{
		PointerIntType3{
			PointerIntField1: &int1,
			PointerIntField2: nil,
			PointerIntField3: &int3,
		}, nil},
	{
		PointerIntType3{
			PointerIntField1: nil,
			PointerIntField2: nil,
			PointerIntField3: nil,
		}, nil},
	{
		PointerIntType3{
			PointerIntField1: nil,
			PointerIntField2: &int2,
			PointerIntField3: nil,
		}, nil},
}

//go:generate pregogen -type=PointerIntArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerIntArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerIntArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerIntArrayType -file=$GOFILE -gen=plus

type PointerIntArrayType struct {
	PointerIntField []*int `json:"pointerintfield"`
}

var intarray1 []*int = []*int{&int1, nil, &int3}
var intarray2 []*int = []*int{nil, nil, nil}

// representative example of data stored in target application
var PointerIntArrayType_examples = []struct {
	PointerIntArrayType
	want []byte
}{
	{
		PointerIntArrayType{
			PointerIntField: intarray1,
		}, nil},
	{
		PointerIntArrayType{
			PointerIntField: intarray2,
		}, nil},
}

//go:generate pregogen -type=PointerIntArrayType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerIntArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerIntArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerIntArrayType3 -file=$GOFILE -gen=plus
type PointerIntArrayType3 struct {
	PointerIntArrayField1 []*int `json:"pointerintfield1"`
	PointerIntArrayField2 []*int `json:"pointerintfield2"`
	PointerIntArrayField3 []*int `json:"pointerintfield3"`
}

var intarray3 []*int = []*int{nil, &int3, nil}

// representative example(s) of data stored in target application
var PointerIntArrayType3_examples = []struct {
	PointerIntArrayType3
	want []byte
}{
	{
		PointerIntArrayType3{
			PointerIntArrayField1: intarray1,
			PointerIntArrayField2: nil,
			PointerIntArrayField3: intarray3,
		}, nil},
	{
		PointerIntArrayType3{
			PointerIntArrayField1: nil,
			PointerIntArrayField2: nil,
			PointerIntArrayField3: nil,
		}, nil},
	{
		PointerIntArrayType3{
			PointerIntArrayField1: nil,
			PointerIntArrayField2: intarray2,
			PointerIntArrayField3: nil,
		}, nil},
}
