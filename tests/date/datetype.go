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
