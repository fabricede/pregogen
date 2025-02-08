//go:generate pregogen -type=DateType -file=$GOFILE -gen=test
//go:generate pregogen -type=DateType -file=$GOFILE -gen=append
//go:generate pregogen -type=DateType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=DateType -file=$GOFILE -gen=plus

package datetype

import "time"

var testTime = time.Now()

type DateType struct {
	DateField time.Time `json:"datefield"`
}

var DateFieldFormat string = time.RFC3339Nano

// representative example of data stored in target application
var DateType_examples = []struct {
	DateType
	want []byte
}{
	{
		DateType{
			DateField: testTime,
		}, nil},
}

//go:generate pregogen -type=DateType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=DateType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=DateType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=DateType3 -file=$GOFILE -gen=plus
type DateType3 struct {
	DateField1 time.Time `json:"datefield1"`
	DateField2 time.Time `json:"datefield2"`
	DateField3 time.Time `json:"datefield3"`
}

var DateField1Format string = time.RFC3339Nano
var DateField2Format string = time.RFC3339Nano
var DateField3Format string = time.RFC3339Nano

// representative example(s) of data stored in target application
var DateType3_examples = []struct {
	DateType3
	want []byte
}{
	{
		DateType3{
			DateField1: testTime,
			DateField2: testTime,
			DateField3: testTime,
		}, nil},
}

//go:generate pregogen -type=DateArrayType -file=$GOFILE -gen=test
//go:generate pregogen -type=DateArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=DateArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=DateArrayType -file=$GOFILE -gen=plus
type DateArrayType struct {
	DateArrayField []time.Time `json:"datearrayfield"`
}

var DateArrayFieldFormat string = time.RFC3339Nano

// representative example of data stored in target application
var DateArrayType_examples = []struct {
	DateArrayType DateArrayType
	want          []byte
}{
	{
		DateArrayType{
			DateArrayField: []time.Time{testTime, testTime, testTime},
		}, nil},
}

type DateArrayType3 struct {
	DateArrayField1 []time.Time `json:"datearray1field"`
	DateArrayField2 []time.Time `json:"datearray2field"`
	DateArrayField3 []time.Time `json:"datearray3field"`
}

var DateArrayField1Format string = time.RFC3339Nano
var DateArrayField2Format string = time.RFC3339Nano
var DateArrayField3Format string = time.RFC3339Nano

//go:generate pregogen -type=DateArrayType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=DateArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=DateArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=DateArrayType3 -file=$GOFILE -gen=plus

// representative example(s) of data stored in target application
var DateArrayType3_examples = []struct {
	DateArrayType3
	want []byte
}{
	{
		DateArrayType3{
			DateArrayField1: []time.Time{testTime, testTime, testTime},
			DateArrayField2: []time.Time{testTime, testTime, testTime},
			DateArrayField3: []time.Time{testTime, testTime, testTime},
		}, nil},
}

//go:generate pregogen -type=PointerDateType -file=$GOFILE -gen=test
//go:generate pregogen -type=PointerDateType -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerDateType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerDateType -file=$GOFILE -gen=plus

type PointerDateType struct {
	PointerDateField *time.Time `json:"pointerdatefield"`
}

var PointerDateFieldFormat string = time.RFC3339Nano
var date1 = testTime

// representative example of data stored in target application
var PointerDateType_examples = []struct {
	PointerDateType
	want []byte
}{
	{
		PointerDateType{
			PointerDateField: &date1,
		}, nil},
	{
		PointerDateType{
			PointerDateField: nil,
		}, nil},
}

//go:generate pregogen -type=PointerDateType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=PointerDateType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerDateType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerDateType3 -file=$GOFILE -gen=plus
type PointerDateType3 struct {
	PointerDateField1 *time.Time `json:"pointerdatefield1"`
	PointerDateField2 *time.Time `json:"pointerdatefield2"`
	PointerDateField3 *time.Time `json:"pointerdatefield3"`
}

var PointerDateField1Format string = time.RFC3339Nano
var PointerDateField2Format string = time.RFC3339Nano
var PointerDateField3Format string = time.RFC3339Nano

var date2 = testTime
var date3 = testTime

// representative example(s) of data stored in target application
var PointerDateType3_examples = []struct {
	PointerDateType3
	want []byte
}{
	{
		PointerDateType3{
			PointerDateField1: &date1,
			PointerDateField2: nil,
			PointerDateField3: &date3,
		}, nil},
	{
		PointerDateType3{
			PointerDateField1: nil,
			PointerDateField2: nil,
			PointerDateField3: nil,
		}, nil},
	{
		PointerDateType3{
			PointerDateField1: nil,
			PointerDateField2: &date2,
			PointerDateField3: nil,
		}, nil},
}

//go:generate pregogen -type=PointerDateArrayType -file=$GOFILE -gen=test
//go:generate pregogen -type=PointerDateArrayType -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerDateArrayType -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerDateArrayType -file=$GOFILE -gen=plus

type PointerDateArrayType struct {
	PointerDateArrayField []*time.Time `json:"pointerdatearrayfield"`
}

var PointerDateArrayFieldFormat string = time.RFC3339Nano

var bytearray0 []*time.Time = []*time.Time{&date1, &date2, &date3}
var bytearray1 []*time.Time = []*time.Time{&date1, nil, &date3}
var bytearray2 []*time.Time = []*time.Time{nil, nil, nil}
var bytearray3 []*time.Time = []*time.Time{nil, &date3, nil}
var bytearray4 []*time.Time = []*time.Time{nil}
var bytearray5 []*time.Time = []*time.Time{&date1}

// representative example of data stored in target application
var PointerDateArrayType_examples = []struct {
	PointerDateArrayType
	want []byte
}{
	{
		PointerDateArrayType{
			PointerDateArrayField: bytearray0,
		}, nil},
	{
		PointerDateArrayType{
			PointerDateArrayField: bytearray1,
		}, nil},
	{
		PointerDateArrayType{
			PointerDateArrayField: bytearray2,
		}, nil},
	{
		PointerDateArrayType{
			PointerDateArrayField: nil,
		}, nil},
	{
		PointerDateArrayType{
			PointerDateArrayField: bytearray3,
		}, nil},
	{
		PointerDateArrayType{
			PointerDateArrayField: bytearray4,
		}, nil},
	{
		PointerDateArrayType{
			PointerDateArrayField: bytearray5,
		}, nil},
}

//go:generate pregogen -type=PointerDateArrayType3 -file=$GOFILE -gen=test
//go:generate pregogen -type=PointerDateArrayType3 -file=$GOFILE -gen=append
//go:generate pregogen -type=PointerDateArrayType3 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=PointerDateArrayType3 -file=$GOFILE -gen=plus
type PointerDateArrayType3 struct {
	PointerDateArrayField1 []*time.Time `json:"pointerdatearrayfield1"`
	PointerDateArrayField2 []*time.Time `json:"pointerdatearrayfield2"`
	PointerDateArrayField3 []*time.Time `json:"pointerdatearrayfield3"`
}

var PointerDateArrayField1Format string = time.RFC3339Nano
var PointerDateArrayField2Format string = time.RFC3339Nano
var PointerDateArrayField3Format string = time.RFC3339Nano

// representative example(s) of data stored in target application
var PointerDateArrayType3_examples = []struct {
	PointerDateArrayType3
	want []byte
}{
	{
		PointerDateArrayType3{
			PointerDateArrayField1: bytearray1,
			PointerDateArrayField2: nil,
			PointerDateArrayField3: bytearray3,
		}, nil},
	{
		PointerDateArrayType3{
			PointerDateArrayField1: nil,
			PointerDateArrayField2: nil,
			PointerDateArrayField3: nil,
		}, nil},
	{
		PointerDateArrayType3{
			PointerDateArrayField1: nil,
			PointerDateArrayField2: bytearray2,
			PointerDateArrayField3: nil,
		}, nil},
}
