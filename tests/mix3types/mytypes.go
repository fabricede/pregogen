//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=testMarshal
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=append
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=plus
package mix3types

import (
	"strings"
	"time"
)

var testTime = time.Now()

// var convertTime, err = time.Parse(time.RFC3339, DateFieldFormat)

type YourType2 struct {
	RuneField  rune      `json:"runefield"`
	RunesField []rune    `json:"runesfield"`
	DateField  time.Time `json:"datefield"`
}

// representative example of data stored in target application (size used for capacity)
var YourType2_examples = []struct {
	YourType2
	want []byte
}{
	{
		YourType2{
			RuneField:  'b',
			RunesField: []rune("123"),
			DateField:  testTime,
		}, nil},
}

var DateFieldFormat string = time.RFC3339Nano

// UnmarshalJSON implements the json.Unmarshaler interface.
func (y *YourType2) UnmarshalJSON_pathfinder(data []byte) (err error) {
	// Custom unmarshaling logic here
	sdata := string(data)
	lendata := len(sdata)

	// Define a type that holds both key and Field info.
	type OrderedField struct {
		Key   string
		start int
		begin int
	}

	var orderedFields []OrderedField

	// List the keys you expect. You can adjust the order or use any iteration order.
	for _, key := range []string{"runefield", "runesfield", "datefield"} {
		posField := strings.Index(sdata, key)
		if posField != -1 {
			// Found the field: store its position info.
			newOf := OrderedField{
				Key:   key,
				start: posField - 2,
				begin: strings.Index(sdata[posField:], ":") + posField + 1,
			}
			inserted := false
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
			}
		}
	}

	for i, field := range orderedFields {
		dataEnd := lendata - 1
		if len(orderedFields) > i+1 {
			dataEnd = orderedFields[i+1].start - 1
		}
		switch field.Key {
		case "runefield":
			y.RuneField = rune(sdata[field.begin+1])
		case "runesfield":
			y.RunesField = []rune(sdata[field.begin+1 : dataEnd])
		case "datefield":
			df, e := time.Parse(DateFieldFormat, sdata[field.begin+1:dataEnd-1])
			if e != nil {
				err = e
			} else {
				y.DateField = df
			}
		}
	}
	return err
}
