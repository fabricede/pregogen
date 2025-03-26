package inttype

import (
	"strconv"
	"testing"
)

var testCases = []struct {
	input    string
	expected int
}{
	{"123", 123},
	{"-456", -456},
	{"0", 0},
	{"999999", 999999},
	{"-1000000", -1000000},
}

func BenchmarkNumberParsing(b *testing.B) {
	b.Run("parseNumber", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				num, _, err := parseNumber([]byte(tc.input))
				if err != nil {
					b.Fatal(err)
				}
				if num != tc.expected {
					b.Fatalf("expected %d, got %d", tc.expected, num)
				}
			}
		}
	})

	b.Run("strconv.ParseInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, tc := range testCases {
				num, err := strconv.ParseInt(tc.input, 10, 32)
				if err != nil {
					b.Fatal(err)
				}
				if int(num) != tc.expected {
					b.Fatalf("expected %d, got %d", tc.expected, num)
				}
			}
		}
	})
}

/* bench results (with new parseintbis method)

BenchmarkIntType_UnmarshalJSON/Test1_UnmarshalJSON-8                       	13792291	        85.55 ns/op	      32 B/op	       1 allocs/op
BenchmarkIntType_UnmarshalJSON/Test1_json.Unmarshal-8                      	 2832322	       461.6 ns/op	     184 B/op	       3 allocs/op
BenchmarkIntType_UnmarshalJSON/Test1_json.Unmarshal_ref-8                  	 2008407	       599.7 ns/op	     216 B/op	       4 allocs/op

BenchmarkIntType3_UnmarshalJSON/Test1_UnmarshalJSON-8                      	 3414820	       351.8 ns/op	     288 B/op	       4 allocs/op
BenchmarkIntType3_UnmarshalJSON/Test1_json.Unmarshal-8                     	 1154444	      1071 ns/op	     440 B/op	       6 allocs/op
BenchmarkIntType3_UnmarshalJSON/Test1_json.Unmarshal_ref-8                 	 1007752	      1176 ns/op	     216 B/op	       4 allocs/op

BenchmarkPointerIntType_UnmarshalJSON/Test1_UnmarshalJSON-8                          	 9744850	       119.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkPointerIntType_UnmarshalJSON/Test1_json.Unmarshal-8                         	 2351108	       530.5 ns/op	     192 B/op	       4 allocs/op
BenchmarkPointerIntType_UnmarshalJSON/Test1_json.Unmarshal_ref-8                     	 1857154	       638.5 ns/op	     216 B/op	       4 allocs/op

BenchmarkPointerIntType3_UnmarshalJSON/Test1_UnmarshalJSON-8         	 2390757	       442.6 ns/op	     320 B/op	       6 allocs/op
BenchmarkPointerIntType3_UnmarshalJSON/Test1_json.Unmarshal-8        	  844656	      1483 ns/op	     472 B/op	       8 allocs/op
BenchmarkPointerIntType3_UnmarshalJSON/Test1_json.Unmarshal_ref-8    	  899997	      1153 ns/op	     216 B/op	       4 allocs/op

*/

/* bench results (with old parseint method)

BenchmarkIntType_UnmarshalJSON/Test1_UnmarshalJSON-8         	11007603	       109.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkIntType_UnmarshalJSON/Test1_json.Unmarshal-8        	 2500605	       478.1 ns/op	     184 B/op	       3 allocs/op
BenchmarkIntType_UnmarshalJSON/Test1_json.Unmarshal_ref-8    	 1937671	       600.1 ns/op	     216 B/op	       4 allocs/op

BenchmarkIntType3_UnmarshalJSON/Test1_UnmarshalJSON-8         	 3055796	       365.3 ns/op	     288 B/op	       4 allocs/op
BenchmarkIntType3_UnmarshalJSON/Test1_json.Unmarshal-8        	 1122579	      1035 ns/op	     440 B/op	       6 allocs/op
BenchmarkIntType3_UnmarshalJSON/Test1_json.Unmarshal_ref-8    	  880150	      1207 ns/op	     216 B/op	       4 allocs/op

BenchmarkPointerIntType_UnmarshalJSON/Test1_UnmarshalJSON-8         	 7701824	       136.8 ns/op	      40 B/op	       2 allocs/op
BenchmarkPointerIntType_UnmarshalJSON/Test1_json.Unmarshal-8        	 2353264	       521.3 ns/op	     192 B/op	       4 allocs/op
BenchmarkPointerIntType_UnmarshalJSON/Test1_json.Unmarshal_ref-8    	 1984062	       661.1 ns/op	     216 B/op	       4 allocs/op

BenchmarkPointerIntType3_UnmarshalJSON/Test1_UnmarshalJSON-8         	 2423326	       463.7 ns/op	     320 B/op	       6 allocs/op
BenchmarkPointerIntType3_UnmarshalJSON/Test1_json.Unmarshal-8        	  709141	      1529 ns/op	     472 B/op	       8 allocs/op
BenchmarkPointerIntType3_UnmarshalJSON/Test1_json.Unmarshal_ref-8    	  886952	      1300 ns/op	     216 B/op	       4 allocs/op

*/
