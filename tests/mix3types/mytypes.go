//go:generate pregogen -type=YourType2 -file=mytypes.go -gen=append
package mix3types

type Mix5Types struct {
	BoolField    bool    `json:"boolfield"`
	Float64Field float64 `json:"float64field"`
	Int8Field    int8    `json:"int8field"`
	IntField     int     `json:"intfield"`
	StringField  string  `json:"stringfield"`
}

// representative example of data stored in target application
var YourType1_example = Mix5Types{
	StringField: "hello", IntField: 123, Int8Field: 12, BoolField: true, Float64Field: 12.34,
}

type YourType2 struct {
	RuneField    rune    `json:"runefield"`
	RunesField   []rune  `json:"runesfield"`
	Float32Field float32 `json:"float32field"`
}

// representative example of data stored in target application
var YourType2_example = YourType2{
	RuneField: 'b', RunesField: []rune("123"), Float32Field: 12.34,
}
