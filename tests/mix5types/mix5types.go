//go:generate pregogen -type=YourType1 -file=mix5types.go -gen=test
//go:generate pregogen -type=YourType1 -file=mix5types.go -gen=append
//go:generate pregogen -type=YourType1 -file=mix5types.go -gen=bytesBuffer
//go:generate pregogen -type=YourType1 -file=mix5types.go -gen=plus
package mypac

type YourType1 struct {
	BoolField    bool    `json:"boolfield"`
	Float64Field float64 `json:"float64field"`
	Int8Field    int8    `json:"int8field"`
	IntField     int     `json:"intfield"`
	StringField  string  `json:"stringfield"`
}

// representative example of data stored in target application (size used for capacity)
var YourType1_examples = []struct {
	YourType1
	want []byte
}{
	{
		YourType1{
			StringField:  "hello",
			IntField:     123,
			Int8Field:    12,
			BoolField:    true,
			Float64Field: 12.34,
		}, []byte(`{"boolfield":true,"float64field":12.34,"int8field":12,"intfield":123,"stringfield":"hello"}`)},
}
