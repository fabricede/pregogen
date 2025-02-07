//go:generate pregogen -type=ByteType -file=$GOFILE -gen=test
//go:generate pregogen -type=ByteType -file=$GOFILE -gen=append
//go:generate pregogen -type=ByteType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=ByteType -file=$GOFILE -gen=plus

package bytetype

type ByteType struct {
	ByteField byte `json:"bytefield"`
}

// representative example of data stored in target application
var ByteType_examples = []struct {
	ByteType
	want []byte
}{
	{
		ByteType{
			ByteField: 0,
		}, nil},
	{
		ByteType{
			ByteField: 1,
		}, nil},
	{
		ByteType{
			ByteField: 2,
		}, nil},
	{
		ByteType{
			ByteField: 3,
		}, nil},
	{
		ByteType{
			ByteField: 4,
		}, nil},
	{
		ByteType{
			ByteField: 5,
		}, nil},
	{
		ByteType{
			ByteField: 6,
		}, nil},
	{
		ByteType{
			ByteField: 7,
		}, nil},
	{
		ByteType{
			ByteField: 8,
		}, nil},
	{
		ByteType{
			ByteField: 9,
		}, nil},
	{
		ByteType{
			ByteField: 10,
		}, nil},
	{
		ByteType{
			ByteField: 11,
		}, nil},
	{
		ByteType{
			ByteField: 12,
		}, nil},
	{
		ByteType{
			ByteField: 13,
		}, nil},
	{
		ByteType{
			ByteField: 14,
		}, nil},
	{
		ByteType{
			ByteField: 15,
		}, nil},
	{
		ByteType{
			ByteField: 16,
		}, nil},
	{
		ByteType{
			ByteField: 17,
		}, nil},
	{
		ByteType{
			ByteField: 18,
		}, nil},
	{
		ByteType{
			ByteField: 19,
		}, nil},
	{
		ByteType{
			ByteField: 20,
		}, nil},
	{
		ByteType{
			ByteField: 21,
		}, nil},
	{
		ByteType{
			ByteField: 22,
		}, nil},
	{
		ByteType{
			ByteField: 23,
		}, nil},
	{
		ByteType{
			ByteField: 24,
		}, nil},
	{
		ByteType{
			ByteField: 25,
		}, nil},
	{
		ByteType{
			ByteField: 26,
		}, nil},
	{
		ByteType{
			ByteField: 27,
		}, nil},
	{
		ByteType{
			ByteField: 28,
		}, nil},
	{
		ByteType{
			ByteField: 29,
		}, nil},
	{
		ByteType{
			ByteField: 30,
		}, nil},
	{
		ByteType{
			ByteField: 31,
		}, nil},
	{
		ByteType{
			ByteField: 32,
		}, nil},
	{
		ByteType{
			ByteField: 33,
		}, nil},
	{
		ByteType{
			ByteField: 34,
		}, nil},
	{
		ByteType{
			ByteField: 35,
		}, nil},
	{
		ByteType{
			ByteField: 36,
		}, nil},
	{
		ByteType{
			ByteField: 37,
		}, nil},
	{
		ByteType{
			ByteField: 38,
		}, nil},
	{
		ByteType{
			ByteField: 39,
		}, nil},
	{
		ByteType{
			ByteField: 40,
		}, nil},
	{
		ByteType{
			ByteField: 41,
		}, nil},
	{
		ByteType{
			ByteField: 42,
		}, nil},
	{
		ByteType{
			ByteField: 43,
		}, nil},
	{
		ByteType{
			ByteField: 44,
		}, nil},
	{
		ByteType{
			ByteField: 45,
		}, nil},
	{
		ByteType{
			ByteField: 46,
		}, nil},
	{
		ByteType{
			ByteField: 47,
		}, nil},
	{
		ByteType{
			ByteField: 48,
		}, nil},
	{
		ByteType{
			ByteField: 49,
		}, nil},
	{
		ByteType{
			ByteField: 50,
		}, nil},
	{
		ByteType{
			ByteField: 51,
		}, nil},
	{
		ByteType{
			ByteField: 52,
		}, nil},
	{
		ByteType{
			ByteField: 53,
		}, nil},
	{
		ByteType{
			ByteField: 54,
		}, nil},
	{
		ByteType{
			ByteField: 55,
		}, nil},
	{
		ByteType{
			ByteField: 56,
		}, nil},
	{
		ByteType{
			ByteField: 57,
		}, nil},
	{
		ByteType{
			ByteField: 58,
		}, nil},
	{
		ByteType{
			ByteField: 59,
		}, nil},
	{
		ByteType{
			ByteField: 60,
		}, nil},
	{
		ByteType{
			ByteField: 61,
		}, nil},
	{
		ByteType{
			ByteField: 62,
		}, nil},
	{
		ByteType{
			ByteField: 63,
		}, nil},
	{
		ByteType{
			ByteField: 64,
		}, nil},
	{
		ByteType{
			ByteField: 65,
		}, nil},
	{
		ByteType{
			ByteField: 66,
		}, nil},
	{
		ByteType{
			ByteField: 67,
		}, nil},
	{
		ByteType{
			ByteField: 68,
		}, nil},
	{
		ByteType{
			ByteField: 69,
		}, nil},
	{
		ByteType{
			ByteField: 70,
		}, nil},
	{
		ByteType{
			ByteField: 71,
		}, nil},
	{
		ByteType{
			ByteField: 72,
		}, nil},
	{
		ByteType{
			ByteField: 73,
		}, nil},
	{
		ByteType{
			ByteField: 74,
		}, nil},
	{
		ByteType{
			ByteField: 75,
		}, nil},
	{
		ByteType{
			ByteField: 76,
		}, nil},
	{
		ByteType{
			ByteField: 77,
		}, nil},
	{
		ByteType{
			ByteField: 78,
		}, nil},
	{
		ByteType{
			ByteField: 79,
		}, nil},
	{
		ByteType{
			ByteField: 255,
		}, nil},
}

//go:generate pregogen -type=ByteType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=ByteType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=ByteType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=ByteType3 -file=$GOFILE -gen=plus
type ByteType3 struct {
	ByteField1 byte `json:"bytefield1"`
	ByteField2 byte `json:"bytefield2"`
	ByteField3 byte `json:"bytefield3"`
}

// representative example(s) of data stored in target application
var ByteType3_examples = []struct {
	ByteType3
	want []byte
}{
	{
		ByteType3{
			ByteField1: 123,
			ByteField2: 67,
			ByteField3: 10,
		}, nil},
}

//go:generate pregogen -type=ByteArrayType -file=$GOFILE -gen=test
//go:generate pregogen -type=ByteArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=ByteArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=ByteArrayType -file=$GOFILE -gen=plus
type ByteArrayType struct {
	ByteArrayField []byte `json:"bytearrayfield"`
}

// representative example of data stored in target application
var ByteArrayType_examples = []struct {
	ByteArrayType ByteArrayType
	want          []byte
}{
	{
		ByteArrayType{
			ByteArrayField: []byte{1, 2, 3},
		}, nil},
}

type ByteArrayType3 struct {
	ByteArrayField1 []byte `json:"bytearray1field"`
	ByteArrayField2 []byte `json:"bytearray2field"`
	ByteArrayField3 []byte `json:"bytearray3field"`
}

//go:generate pregogen -type=ByteArrayType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=ByteArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=ByteArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=ByteArrayType3 -file=$GOFILE -gen=plus

// representative example(s) of data stored in target application
var ByteArrayType3_examples = []struct {
	ByteArrayType3
	want []byte
}{
	{
		ByteArrayType3{
			ByteArrayField1: []byte{1, 2, 3},
			ByteArrayField2: []byte{4, 5, 6},
			ByteArrayField3: []byte{7, 8, 9},
		}, nil},
}

//go:generate pregogen -type=PointerByteType -file=$GOFILE -gen=test
//go:generate pregogen -type=PointerByteType -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerByteType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerByteType -file=$GOFILE -gen=plus

type PointerByteType struct {
	PointerByteField *byte `json:"pointerbytefield"`
}

var byte1 byte = 64

// representative example of data stored in target application
var PointerByteType_examples = []struct {
	PointerByteType
	want []byte
}{
	{
		PointerByteType{
			PointerByteField: &byte1,
		}, nil},
	{
		PointerByteType{
			PointerByteField: nil,
		}, nil},
}

//go:generate pregogen -type=PointerByteType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=PointerByteType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerByteType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerByteType3 -file=$GOFILE -gen=plus
type PointerByteType3 struct {
	PointerByteField1 *byte `json:"pointerbytefield1"`
	PointerByteField2 *byte `json:"pointerbytefield2"`
	PointerByteField3 *byte `json:"pointerbytefield3"`
}

var byte2 byte = 125
var byte3 byte = 90

// representative example(s) of data stored in target application
var PointerByteType3_examples = []struct {
	PointerByteType3
	want []byte
}{
	{
		PointerByteType3{
			PointerByteField1: &byte1,
			PointerByteField2: nil,
			PointerByteField3: &byte3,
		}, nil},
	{
		PointerByteType3{
			PointerByteField1: nil,
			PointerByteField2: nil,
			PointerByteField3: nil,
		}, nil},
	{
		PointerByteType3{
			PointerByteField1: nil,
			PointerByteField2: &byte2,
			PointerByteField3: nil,
		}, nil},
}
