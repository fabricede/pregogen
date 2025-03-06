//go:generate pregogen -type=StringType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=StringType -file=$GOFILE -gen=marshal
//go:generate pregogen -type=StringType -file=$GOFILE -gen=unmarshal
package stringtype

import "time"

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

//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=testAll
//go:generate pregogen -type=StringType3 -file=$GOFILE -gen=marshal
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
//go:generate pregogen -type=StringArrayType -file=$GOFILE -gen=marshal
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
//go:generate pregogen -type=StringArrayType3 -file=$GOFILE -gen=marshal

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

//go:generate pregogen -type=PointerStringType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=PointerStringType -file=$GOFILE -gen=marshal
//go:generate pregogen -type=PointerStringType -file=$GOFILE -gen=unmarshal
type PointerStringType struct {
	PointerStringField *string `json:"pointerstringfield"`
}

var string1 = "toto"

// representative example of data stored in target application
var PointerStringType_examples = []struct {
	PointerStringType
	want []byte
}{
	{
		PointerStringType{
			PointerStringField: &string1,
		}, nil},
	{
		PointerStringType{
			PointerStringField: nil,
		}, nil},
}

//go:generate pregogen -type=PointerStringType3 -file=$GOFILE -gen=testAll
//go:generate pregogen -type=PointerStringType3 -file=$GOFILE -gen=marshal
//go:generate pregogen -type=PointerStringType3 -file=$GOFILE -gen=unmarshal
type PointerStringType3 struct {
	PointerStringField1 *string `json:"pointerstringfield1"`
	PointerStringField2 *string `json:"pointerstringfield2"`
	PointerStringField3 *string `json:"pointerstringfield3"`
}

var string2 = "tata"
var string3 = "titi"

// representative example(s) of data stored in target application
var PointerStringType3_examples = []struct {
	PointerStringType3
	want []byte
}{
	{
		PointerStringType3{
			PointerStringField1: &string1,
			PointerStringField2: nil,
			PointerStringField3: &string3,
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
			PointerStringField2: &string2,
			PointerStringField3: nil,
		}, nil},
}

//go:generate pregogen -type=PointerStringArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerStringArrayType -file=$GOFILE -gen=marshal
type PointerStringArrayType struct {
	PointerStringField []*string `json:"pointerstringfield"`
}

var stringarray1 []*string = []*string{&string1, nil, &string2}
var stringarray2 []*string = []*string{nil, nil, nil}

// representative example of data stored in target application
var PointerStringArrayType_examples = []struct {
	PointerStringArrayType
	want []byte
}{
	{
		PointerStringArrayType{
			PointerStringField: stringarray1,
		}, nil},
	{
		PointerStringArrayType{
			PointerStringField: stringarray2,
		}, nil},
}

//go:generate pregogen -type=PointerStringArrayType3 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerStringArrayType3 -file=$GOFILE -gen=marshal
type PointerStringArrayType3 struct {
	PointerStringArrayField1 []*string `json:"pointerstringfield1"`
	PointerStringArrayField2 []*string `json:"pointerstringfield2"`
	PointerStringArrayField3 []*string `json:"pointerstringfield3"`
}

var stringarray3 []*string = []*string{nil, &string3, nil}

// representative example(s) of data stored in target application
var PointerStringArrayType3_examples = []struct {
	PointerStringArrayType3
	want []byte
}{
	{
		PointerStringArrayType3{
			PointerStringArrayField1: stringarray1,
			PointerStringArrayField2: nil,
			PointerStringArrayField3: stringarray3,
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
			PointerStringArrayField2: stringarray2,
			PointerStringArrayField3: nil,
		}, nil},
}

//go:generate pregogen -type=DateType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=DateType -file=$GOFILE -gen=marshal
//go:generate pregogen -type=DateType -file=$GOFILE -gen=unmarshal

// Date is a time.Time in Go, but is a "string" type in terms of JSON
type DateType struct {
	DateField time.Time `json:"datefield"`
}

// DateFieldFormat is the format used to marshal/unmarshal the DateField
var DateFieldFormat string = time.RFC3339Nano

// testTime is a fixed time used for tests (use compiles time)
// Round(0) is used to remove the monotonic clock part of the time
var testTime = time.Now().Round(0)

// representative example of data stored in target application (size used for capacity)
var DateType_examples = []struct {
	DateType
	want []byte
}{
	{
		DateType{

			DateField: testTime,
		}, nil},
}

//----------------------------
// Pointer Date
//----------------------------

// PointerDateType is a date type stored as a pointer.
type PointerDateType struct {
	DateField *time.Time `json:"datefield"`
}

var PointerDateType_examples = []struct {
	PointerDateType
	want []byte
}{
	{
		PointerDateType{
			DateField: &testTime,
		}, nil,
	},
	{
		PointerDateType{
			DateField: nil,
		}, nil,
	},
}

//----------------------------
// Array of Dates
//----------------------------

var DateFieldsFormat string = time.RFC3339Nano

// DateArrayType is a slice of dates.
type DateArrayType struct {
	DateFields []time.Time `json:"datefields"`
}

var DateArrayType_examples = []struct {
	DateArrayType
	want []byte
}{
	{
		DateArrayType{
			DateFields: []time.Time{
				testTime.Round(0),
				testTime.Add(time.Hour).Round(0),
			},
		}, nil,
	},
}

//----------------------------
// Array of Pointer Dates
//----------------------------

// PointerDateArrayType is a slice of pointers to dates.
type PointerDateArrayType struct {
	DateFields []*time.Time `json:"datefields"`
}

var PointerDateArrayType_examples = []struct {
	PointerDateArrayType
	want []byte
}{
	{
		PointerDateArrayType{
			DateFields: []*time.Time{&testTime, nil, &testTime},
		}, nil,
	},
}

//go:generate pregogen -type=DateType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=DateType -file=$GOFILE -gen=marshal
//go:generate pregogen -type=DateType -file=$GOFILE -gen=unmarshal

//go:generate pregogen -type=PointerDateType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=PointerDateType -file=$GOFILE -gen=marshal
//go:generate pregogen -type=PointerDateType -file=$GOFILE -gen=unmarshal

//go:generate pregogen -type=DateArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=DateArrayType -file=$GOFILE -gen=marshal

//go:generate pregogen -type=PointerDateArrayType -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=PointerDateArrayType -file=$GOFILE -gen=marshal
