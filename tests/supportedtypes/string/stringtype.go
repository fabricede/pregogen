//go:generate pregogen -type=StringType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=StringType -file=$GOFILE -gen=append
//go:generate pregogen -type=StringType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=StringType -file=$GOFILE -gen=plus
//go:generate pregogen -type=StringType -file=$GOFILE -gen=unmarshal
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

//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=plus
//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=unmarshal
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

//go:generate pregogen -type=StringArrayType -file=$GOFILE -gen=testMarshal
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

//go:generate pregogen -type=StringArrayType3 -file=$GOFILE -gen=testMarshal
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

//go:generate pregogen -type=PointerStringType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerStringType -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerStringType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerStringType -file=$GOFILE -gen=plus

type PointerStringType struct {
	PointerStringField *string `json:"pointerstringfield"`
}

var int1 = "toto"

// representative example of data stored in target application
var PointerStringType_examples = []struct {
	PointerStringType
	want []byte
}{
	{
		PointerStringType{
			PointerStringField: &int1,
		}, nil},
	{
		PointerStringType{
			PointerStringField: nil,
		}, nil},
}

//go:generate pregogen -type=PointerStringType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerStringType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerStringType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerStringType3 -file=$GOFILE -gen=plus
type PointerStringType3 struct {
	PointerStringField1 *string `json:"pointerstringfield1"`
	PointerStringField2 *string `json:"pointerstringfield2"`
	PointerStringField3 *string `json:"pointerstringfield3"`
}

var int2 = "tata"
var int3 = "titi"

// representative example(s) of data stored in target application
var PointerStringType3_examples = []struct {
	PointerStringType3
	want []byte
}{
	{
		PointerStringType3{
			PointerStringField1: &int1,
			PointerStringField2: nil,
			PointerStringField3: &int3,
		}, nil},
	{
		PointerStringType3{
			PointerStringField1: nil,
			PointerStringField2: nil,
			PointerStringField3: nil,
		}, nil},
	{
		PointerStringType3{
			PointerStringField1: nil,
			PointerStringField2: &int2,
			PointerStringField3: nil,
		}, nil},
}

//go:generate pregogen -type=PointerStringArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerStringArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerStringArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerStringArrayType -file=$GOFILE -gen=plus
type PointerStringArrayType struct {
	PointerStringField []*string `json:"pointerstringfield"`
}

var intarray1 []*string = []*string{&int1, nil, &int2}
var intarray2 []*string = []*string{nil, nil, nil}

// representative example of data stored in target application
var PointerStringArrayType_examples = []struct {
	PointerStringArrayType
	want []byte
}{
	{
		PointerStringArrayType{
			PointerStringField: intarray1,
		}, nil},
	{
		PointerStringArrayType{
			PointerStringField: intarray2,
		}, nil},
}

//go:generate pregogen -type=PointerStringArrayType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerStringArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerStringArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerStringArrayType3 -file=$GOFILE -gen=plus
type PointerStringArrayType3 struct {
	PointerStringArrayField1 []*string `json:"pointerstringfield1"`
	PointerStringArrayField2 []*string `json:"pointerstringfield2"`
	PointerStringArrayField3 []*string `json:"pointerstringfield3"`
}

var intarray3 []*string = []*string{nil, &int3, nil}

// representative example(s) of data stored in target application
var PointerStringArrayType3_examples = []struct {
	PointerStringArrayType3
	want []byte
}{
	{
		PointerStringArrayType3{
			PointerStringArrayField1: intarray1,
			PointerStringArrayField2: nil,
			PointerStringArrayField3: intarray3,
		}, nil},
	{
		PointerStringArrayType3{
			PointerStringArrayField1: nil,
			PointerStringArrayField2: nil,
			PointerStringArrayField3: nil,
		}, nil},
	{
		PointerStringArrayType3{
			PointerStringArrayField1: nil,
			PointerStringArrayField2: intarray2,
			PointerStringArrayField3: nil,
		}, nil},
}
