# pregogen

A code generator for custom JSON methods in Go.

## Current State
First commit demonstrates feasibility and increase in json Marshal performance over the standard library.

results with `plus` on singletype:
goos: linux
goarch: amd64
pkg: pregogen/tests/singletype/bool
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkBoolType_MarshalJSON/Test1_MarshalJSON-8         	34644636	        35.24 ns/op	      24 B/op	       1 allocs/op
BenchmarkBoolType_MarshalJSON/Test1_json.Marshal-8        	 3460633	       319.7 ns/op	      48 B/op	       2 allocs/op
BenchmarkBoolType_MarshalJSON/Test1_ref-8                 	 7754834	       148.7 ns/op	      24 B/op	       1 allocs/op

results with `append` on mix5types:
goos: linux
goarch: amd64
pkg: pregogen/tests/mix5types
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkYourType1_MarshalJSON/Test1_MarshalJSON-8         	 2617011	       463.7 ns/op	     195 B/op	       5 allocs/op
BenchmarkYourType1_MarshalJSON/Test1_json.Marshal-8        	  774948	      1428 ns/op	     339 B/op	       7 allocs/op
BenchmarkYourType1_MarshalJSON/Test1_ref-8                 	 2249636	       526.7 ns/
op	     144 B/op	       2 allocs/op

## Overview
- Parses Go source files and generates JSON methods using templates.
- Generates multiple implementations with tests and benchmarks.

## Installation
1. Clone this repository.
2. Run `go build main/pregogen.go`.
3. Add executable to the path : `export PATH="$PATH:$(pwd)"`.

## Usage
- Use the `-type` flag to specify the struct type.
- Use the `-file` flag to specify the Go source file.
- Use the `-gen`  flag to specify implementation method (currently `append`, `bytesBuffer` and `plus`) or `test` to generate the test file
- Example: `pregogen -type MyStruct -file myfile.go -gen append`.

## License
MIT License.