//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=test
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=append
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=plus
package mix3types

import "time"

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
