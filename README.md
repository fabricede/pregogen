# pregogen

A code generator for custom JSON methods in Go.

## Current State
demonstrates feasibility and increase in json Marshal performance over the standard library.
Handles structure with basic types (bool, int, float, string) date, pointer and array

## TODO 
	- pointer array
	- Enum capabilities inspired from bool
	- Unmarshal
	- improve 'plus' marshal method to avoid multiple allocation

## Overview
- Parses Go source files and generates JSON methods using templates.
- Generates multiple implementations with tests and benchmarks.

## Installation
1. Clone this repository.
2. Run `go build main/pregogen.go`.
3. Add executable to the path : `export PATH="$PATH:$(pwd)"`.
4. To generate tests : `go generate ./...`

## Usage
- Use the `-type` flag to specify the struct type.
- Use the `-file` flag to specify the Go source file.
- Use the `-gen`  flag to specify implementation method (currently `append`, `bytesBuffer` and `plus`) or `test` to generate the test file
- Example: `pregogen -type MyStruct -file myfile.go -gen append`.

## License
MIT License.

## Benchmarks
results on bool:
goos: linux
goarch: amd64
pkg: pregogen/tests/basictype/bool
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkBoolArrayType3_MarshalJSON/Test1_MarshalJSON_bytesBuffer-8         	 4633483	       220.0 ns/op	     192 B/op	       2 allocs/op
BenchmarkBoolArrayType3_MarshalJSON/Test1_MarshalJSON_plus-8                	 2228943	       604.8 ns/op	     736 B/op	       9 allocs/op
BenchmarkBoolArrayType3_MarshalJSON/Test1_MarshalJSON_append-8              	 4715997	       226.7 ns/op	     360 B/op	       4 allocs/op
BenchmarkBoolArrayType3_MarshalJSON/Test1_json.Marshal_bytesBuffer-8        	 2490760	       470.4 ns/op	     192 B/op	       2 allocs/op
BenchmarkBoolArrayType3_MarshalJSON/Test1_ref-8                             	 2681134	       488.9 ns/op	     192 B/op	       2 allocs/op
2025/01/25 19:33:20 a_bytesBuffer:  {"boolarray1field":[true,true,false],"boolarray2field":[false,true,true],"boolarray3field":[true,false,false]}  a_plus:  {"boolarray1field":[true,true,false],"boolarray2field":[false,true,true],"boolarray3field":[true,false,false]}  a_append : {"boolarray1field":[true,true,false],"boolarray2field":[false,true,true],"boolarray3field":[true,false,false]}
BenchmarkBoolArrayType_MarshalJSON/Test1_MarshalJSON_bytesBuffer-8          	11470828	        87.50 ns/op	      64 B/op	       1 allocs/op
BenchmarkBoolArrayType_MarshalJSON/Test1_MarshalJSON_plus-8                 	 7203508	       154.2 ns/op	     128 B/op	       3 allocs/op
BenchmarkBoolArrayType_MarshalJSON/Test1_MarshalJSON_append-8               	13528222	        90.35 ns/op	      72 B/op	       2 allocs/op
BenchmarkBoolArrayType_MarshalJSON/Test1_json.Marshal_bytesBuffer-8         	 4654036	       266.6 ns/op	      72 B/op	       2 allocs/op
BenchmarkBoolArrayType_MarshalJSON/Test1_ref-8                              	 4267012	       262.6 ns/op	      72 B/op	       2 allocs/op
2025/01/25 19:33:26 a_bytesBuffer:  {"boolarrayfield":[true,true,false]}  a_plus:  {"boolarrayfield":[true,true,false]}  a_append : {"boolarrayfield":[true,true,false]}
BenchmarkBoolType3_MarshalJSON/Test1_MarshalJSON_bytesBuffer-8              	17547643	        74.19 ns/op	      64 B/op	       1 allocs/op
BenchmarkBoolType3_MarshalJSON/Test1_MarshalJSON_plus-8                     	 7695096	       165.9 ns/op	     176 B/op	       3 allocs/op
BenchmarkBoolType3_MarshalJSON/Test1_MarshalJSON_append-8                   	 7220943	       161.0 ns/op	     168 B/op	       3 allocs/op
BenchmarkBoolType3_MarshalJSON/Test1_json.Marshal_bytesBuffer-8             	 5353892	       232.5 ns/op	      67 B/op	       2 allocs/op
BenchmarkBoolType3_MarshalJSON/Test1_ref-8                                  	 4490480	       236.6 ns/op	      67 B/op	       2 allocs/op
2025/01/25 19:33:35 a_bytesBuffer:  {"boolfield1":true,"boolfield2":true,"boolfield3":false}  a_plus:  {"boolfield1":true,"boolfield2":true,"boolfield3":false}  a_append : {"boolfield1":true,"boolfield2":true,"boolfield3":false}
BenchmarkBoolType_MarshalJSON/Test1_MarshalJSON_bytesBuffer-8               	569869083	         2.119 ns/op	       0 B/op	       0 allocs/op
BenchmarkBoolType_MarshalJSON/Test1_MarshalJSON_plus-8                      	526323762	         2.092 ns/op	       0 B/op	       0 allocs/op
BenchmarkBoolType_MarshalJSON/Test1_MarshalJSON_append-8                    	572374892	         2.165 ns/op	       0 B/op	       0 allocs/op
BenchmarkBoolType_MarshalJSON/Test1_json.Marshal_bytesBuffer-8              	 8166038	       146.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkBoolType_MarshalJSON/Test1_ref-8                                   	 8050928	       152.2 ns/op	      24 B/op	       1 allocs/op
2025/01/25 19:33:42 a_bytesBuffer:  {"boolfield":true}  a_plus:  {"boolfield":true}  a_append : {"boolfield":true}
PASS

results on mix5types:
goos: linux
goarch: amd64
pkg: pregogen/tests/mix5types
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkYourType1_MarshalJSON/Test1_MarshalJSON_bytesBuffer-8         	 3009912	       414.8 ns/op	     219 B/op	       4 allocs/op
BenchmarkYourType1_MarshalJSON/Test1_MarshalJSON_plus-8                	 1539234	       790.4 ns/op	     795 B/op	      12 allocs/op
BenchmarkYourType1_MarshalJSON/Test1_MarshalJSON_append-8              	 3418110	       334.5 ns/op	     195 B/op	       5 allocs/op
BenchmarkYourType1_MarshalJSON/Test1_json.Marshal_bytesBuffer-8        	 2268294	       504.4 ns/op	     144 B/op	       2 allocs/op
BenchmarkYourType1_MarshalJSON/Test1_ref-8                             	 2270050	       508.4 ns/op	     144 B/op	       2 allocs/op
2025/01/25 19:28:42 a_bytesBuffer:  {"boolfield":true,"float64field":12.34,"int8field":12,"intfield":123,"stringfield":"hello"}  a_plus:  {"boolfield":true,"float64field":12.34,"int8field":12,"intfield":123,"stringfield":"hello"}  a_append : {"boolfield":true,"float64field":12.34,"int8field":12,"intfield":123,"stringfield":"hello"}