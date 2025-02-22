# pregogen

A code generator for custom JSON methods in Go.

## Current State
demonstrates feasibility and increase in json Marshal performance over the standard library.
Handles structure with basic types (bool, int, float, string) date, []byte, pointer, array and array of pointer

basic Unmarshal

## TODO 
	- Extend Enum capabilities
	- improve 'plus' marshal method to avoid multiple allocation
	- take into account json flag
	- map types
	- composed types

## Overview
- Parses Go source files and generates JSON methods using templates.
- Generates multiple implementations with tests and benchmarks.

## Installation
1. Clone this repository.
2. Run `go build main/pregogen.go`.
3. Add executable to the path : `export PATH="$PATH:$(pwd)"`.
4. To generate tests : `go generate ./...`
5. To execute tests : `go test ./...`

## Usage
- Use the `-type` flag to specify the struct type.
- Use the `-file` flag to specify the Go source file.
- Use the `-gen`  flag to specify implementation method (currently `append`, `bytesBuffer` and `plus`) or `testMarshal` to generate the test file
- Example: `pregogen -type MyStruct -file myfile.go -gen append`.

## License
MIT License.

## Benchmarks

### Single Bool Type
Performance for a struct with a single boolean field:
| Method | Operations/sec | NS/Op | Bytes/Op | Allocs/Op |
|--------|---------------|--------|-----------|------------|
| bytesBuffer | 569,869,083 | 2.119 | 0 | 0 |
| plus | 526,323,762 | 2.092 | 0 | 0 |
| append | 572,374,892 | 2.165 | 0 | 0 |
| json.Marshal | 8,166,038 | 146.5 | 24 | 1 |

Output: `{"boolfield":true}`

### Three Bool Fields
Performance for a struct with three boolean fields:
| Method | Operations/sec | NS/Op | Bytes/Op | Allocs/Op |
|--------|---------------|--------|-----------|------------|
| bytesBuffer | 17,547,643 | 74.19 | 64 | 1 |
| plus | 7,695,096 | 165.9 | 176 | 3 |
| append | 7,220,943 | 161.0 | 168 | 3 |
| json.Marshal | 5,353,892 | 232.5 | 67 | 2 |

Output: `{"boolfield1":true,"boolfield2":true,"boolfield3":false}`

### Bool Array
Performance for a struct with a single boolean array:
| Method | Operations/sec | NS/Op | Bytes/Op | Allocs/Op |
|--------|---------------|--------|-----------|------------|
| bytesBuffer | 11,470,828 | 87.50 | 64 | 1 |
| plus | 7,203,508 | 154.2 | 128 | 3 |
| append | 13,528,222 | 90.35 | 72 | 2 |
| json.Marshal | 4,654,036 | 266.6 | 72 | 2 |

Output: `{"boolarrayfield":[true,true,false]}`

### Mixed Types
Performance for a struct with multiple types (bool, float64, int8, int, string):
| Method | Operations/sec | NS/Op | Bytes/Op | Allocs/Op |
|--------|---------------|--------|-----------|------------|
| bytesBuffer | 3,009,912 | 414.8 | 219 | 4 |
| plus | 1,539,234 | 790.4 | 795 | 12 |
| append | 3,418,110 | 334.5 | 195 | 5 |
| json.Marshal | 2,268,294 | 504.4 | 144 | 2 |

Output: `{"boolfield":true,"float64field":12.34,"int8field":12,"intfield":123,"stringfield":"hello"}`

### Summary
- bytesBuffer consistently performs best across almost all types
- append method shows good balance of performance and memory usage
- plus method generally uses more memory and allocations
- at least one custom method outperform standard json.Marshal