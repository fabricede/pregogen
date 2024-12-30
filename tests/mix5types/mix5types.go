//go:generate pregogen -type=YourType1 -file=mix5types.go -gen=append
package mypac

type YourType1 struct {
	BoolField    bool    `json:"boolfield"`
	Float64Field float64 `json:"float64field"`
	Int8Field    int8    `json:"int8field"`
	IntField     int     `json:"intfield"`
	StringField  string  `json:"stringfield"`
}

// representative example of data stored in target application
var YourType1_example = YourType1{
	StringField: "hello", IntField: 123, Int8Field: 12, BoolField: true, Float64Field: 12.34,
}

type YourType2 struct {
	ByteField    byte    `json:"bytefield"`
	BytesField   []byte  `json:"bytesfield"`
	Float32Field float32 `json:"float32field"`
}

// representative example of data stored in target application
var YourType2_example = YourType2{
	ByteField: 'b', BytesField: []byte("123"), Float32Field: 12.34,
}
