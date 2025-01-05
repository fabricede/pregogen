//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=test
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=append
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=bytesBuffer
//go:generate pregogen -type=YourType2 -file=$GOFILE -gen=plus
package mix3types

type YourType2 struct {
	RuneField    rune    `json:"runefield"`
	RunesField   []rune  `json:"runesfield"`
	Float32Field float32 `json:"float32field"`
}

// representative example of data stored in target application (size used for capacity)
var YourType2_examples = []struct {
	YourType2
	want []byte
}{
	{
		YourType2{
			RuneField:    'b',
			RunesField:   []rune("123"),
			Float32Field: 12.34,
		}, nil},
}
