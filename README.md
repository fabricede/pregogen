# pregogen

A code generator for custom JSON methods in Go.

## Current State
First commit demonstrates feasibility and a ~10% increase in json Marshal performance over the standard library.

## Overview
- Parses Go source files and generates JSON methods using templates.
- Generates multiple implementations with tests and benchmarks.

## Installation
1. Clone this repository.
2. Run `go install ./...`.

## Usage
- Use the `-type` flag to specify the struct type.
- Use the `-file` flag to specify the Go source file.
- Use the `-gen`  flag to specify implementation method (currently `append`, `bytesBuffer` and `plus`) or `test` to generate the test file
- Example: `pregogen -type MyStruct -file myfile.go -gen append`.

## License
MIT License.