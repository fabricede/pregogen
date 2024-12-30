//go:generate pregogen -type=FloatType -file=floattype.go -gen=plus
package floattype

type FloatType struct {
	FloatField float32 `json:"floatfield"`
}

// representative example of data stored in target application
var FloatType_example = FloatType{
	FloatField: 12345.12,
}
