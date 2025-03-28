//go:generate pregogen -type=IntType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=IntType -file=$GOFILE -gen=marshal
//go:generate pregogen -type=IntType -file=$GOFILE -gen=unmarshal
package inttype

import (
	"errors"
	"strconv"
	"strings"
)

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

//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=testAll
//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=marshal
//go:generate pregogen -type=IntType3 -file=$GOFILE -gen=unmarshal
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
//go:generate pregogen -type=IntArrayType -file=$GOFILE -gen=marshal
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
//go:generate pregogen -type=IntArrayType3 -file=$GOFILE -gen=marshal
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

//go:generate pregogen -type=PointerIntType -file=$GOFILE -gen=testAll
//go:generate pregogen -type=PointerIntType -file=$GOFILE -gen=marshal
//go:generate pregogen -type=PointerIntType -file=$GOFILE -gen=unmarshal
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

//go:generate pregogen -type=PointerIntType3 -file=$GOFILE -gen=testAll
//go:generate pregogen -type=PointerIntType3 -file=$GOFILE -gen=marshal
//go:generate pregogen -type=PointerIntType3 -file=$GOFILE -gen=unmarshal
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
//go:generate pregogen -type=PointerIntArrayType -file=$GOFILE -gen=marshal

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
//go:generate pregogen -type=PointerIntArrayType3 -file=$GOFILE -gen=marshal
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

func (i_receiver *IntType) UnmarshalJSON_opt(data []byte) (err error) {
	// UnmarshalJSON implements the json.Unmarshaler interface.
	sdata := string(data)

	//var orderedFields *OrderedField

	// List expected keys.
	//for _, key := range []string{ "intfield", } {
	posField := strings.Index(sdata, "intfield")
	if posField != -1 {
		lendata := len(sdata)
		// Found the field: store its position info.
		begin := posField + len("intfield") + 2 // skip ':'
		dataEnd := lendata - 2                  // remove final '}'
		dfIntField, eIntField := strconv.ParseInt(sdata[begin:dataEnd+1], 10, 32)
		if eIntField != nil {
			err = eIntField
		} else {
			i_receiver.IntField = int(dfIntField)
		}

		/*inserted := false
		for i, of := range orderedFields {
			if newOf.start < of.start {
				orderedFields = append(orderedFields[:i],
					append([]OrderedField{newOf}, orderedFields[i:]...)...)
				inserted = true
				break
			}
		}
		if !inserted {
			orderedFields = append(orderedFields, newOf)
		}*/
	}
	//}

	/*for i, field := range orderedFields {
			dataEnd := lendata - 1
			if len(orderedFields) > i+1 {
				dataEnd = orderedFields[i+1].start - 1
			}
			switch field.Key {

	        case "intfield":
	    dfIntField, eIntField := strconv.ParseInt(sdata[field.begin:dataEnd+1], 10, 64)
	    if eIntField != nil {
	        err = eIntField
	    } else {
	        i_receiver.IntField = int(dfIntField)
	    }

			}
		}*/
	return err
}

func parseNumber(data []byte) (num int, pos int, err error) {
	isNegative := data[pos] == '-'
	if isNegative {
		pos++
	}

	start := pos
	for pos < len(data) && data[pos] >= '0' && data[pos] <= '9' {
		num = num*10 + int(data[pos]-'0')
		pos++
	}

	if pos == start || (isNegative && pos == start+1) {
		return 0, 0, errors.New("invalid number")
	}
	if isNegative {
		num = -num
	}

	return num, pos, nil
}

func (i_receiver *IntArrayType) UnmarshalJSON_array(data []byte) (err error) {
	// Skip whitespace and check for '['
	pos := 0
	for pos < len(data) && (data[pos] == ' ' || data[pos] == '\t' || data[pos] == '\n' || data[pos] == '\r') {
		pos++
	}
	if pos >= len(data) || data[pos] != '[' {
		return errors.New("expected '['")
	}
	pos++

	var result []int
	for pos < len(data) {
		// Skip whitespace
		for pos < len(data) && (data[pos] == ' ' || data[pos] == '\t' || data[pos] == '\n' || data[pos] == '\r') {
			pos++
		}

		if pos >= len(data) {
			return errors.New("unexpected end of data")
		}

		// Check for end of array
		if data[pos] == ']' {
			i_receiver.IntArrayField = result
			return nil
		}

		// Parse number
		num, newPos, err := parseNumber(data[pos:])
		if err != nil {
			return err
		}
		pos = newPos

		result = append(result, num)

		// Skip whitespace
		for pos < len(data) && (data[pos] == ' ' || data[pos] == '\t' || data[pos] == '\n' || data[pos] == '\r') {
			pos++
		}

		// Expect comma or end of array
		if pos >= len(data) {
			return errors.New("unexpected end of data")
		}
		if data[pos] == ']' {
			i_receiver.IntArrayField = result
			return nil
		}
		if data[pos] != ',' {
			return errors.New("expected ',' or ']'")
		}
		pos++
	}

	return errors.New("unexpected end of data")
}
