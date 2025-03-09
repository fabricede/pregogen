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
- Use the `-gen`  flag to specify implementation method (currently `append`, `bytesBuffer` `stringsBuilder` and `plus`) or `testMarshal` to generate the test file
- Example: `pregogen -type MyStruct -file myfile.go -gen append`.

## License
MIT License.

## Benchmarks

**System:**  
- OS/Arch: linux/amd64  
- Package: pregogen/tests/supportedtypes/bool  
- CPU: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz

---

### Benchmark: Bool Array (Type3)
Performance for a struct with three boolean array fields.

| Method                          | Operations/sec | NS/Op    | Bytes/Op | Allocs/Op |
|---------------------------------|----------------|----------|----------|-----------|
| MarshalJSON_bytesBuffer         | 5,587,066      | 198.8 ns/op    | 192 B/op  | 2        |
| MarshalJSON_stringsBuilder      | 3,986,001      | 268.1 ns/op    | 472 B/op  | 5        |
| **MarshalJSON_append**              | **6,112,934**      | **202.4 ns/op**    | 360 B/op  | 4        |
| MarshalJSON_plus                | 2,542,952      | 489.9 ns/op    | 736 B/op  | 9        |
| json.Marshal                    | 2,699,517      | 465.6 ns/op    | 192 B/op  | 2        |
| **ref**                             | **2,556,445**      | **485.4 ns/op**    | 192 B/op  | 2        |

**Output:**  
`{"boolarray1field":[true,true,false],"boolarray2field":[false,true,true],"boolarray3field":[true,false,false]}`

---

### Benchmark: Bool Array (Single Field)
Performance for a struct with a single boolean array field.

| Method                          | Operations/sec | NS/Op    | Bytes/Op | Allocs/Op |
|---------------------------------|----------------|----------|----------|-----------|
| MarshalJSON_bytesBuffer         | 17,481,135     | 79.59 ns/op     | 64 B/op   | 1        |
| MarshalJSON_stringsBuilder      | 10,216,454     | 123.8 ns/op     | 120 B/op  | 3        |
| **MarshalJSON_append**              | **15,211,228**     | **71.29 ns/op**     | 72 B/op   | 2        |
| MarshalJSON_plus                | 8,100,045      | 184.6 ns/op     | 128 B/op  | 3        |
| json.Marshal                    | 4,159,155      | 260.8 ns/op     | 72 B/op   | 2        |
| **ref**                             | **4,269,824**      | **246.2 ns/op**     | 72 B/op   | 2        |

**Output:**  
`{"boolarrayfield":[true,true,false]}`

---

### Benchmark: Bool (Type3)
Performance for a struct with three boolean fields.

| Method                          | Operations/sec | NS/Op    | Bytes/Op | Allocs/Op |
|---------------------------------|----------------|----------|----------|-----------|
| **MarshalJSON_bytesBuffer**        | **19,748,805**     | **64.36 ns/op**     | 64 B/op   | 1        |
| MarshalJSON_stringsBuilder      | 6,474,884      | 194.2 ns/op     | 232 B/op  | 4        |
| MarshalJSON_append              | 8,926,711      | 112.7 ns/op     | 168 B/op  | 3        |
| MarshalJSON_plus                | 8,232,297      | 158.3 ns/op     | 176 B/op  | 3        |
| json.Marshal                    | 5,176,584      | 219.7 ns/op     | 67 B/op   | 2        |
| **ref**                            | **5,316,712**      | **232.4 ns/op**     | 67 B/op   | 2        |

**Output:**  
`{"boolfield1":true,"boolfield2":true,"boolfield3":false}`

---

### Benchmark: Bool (Single Field)
Performance for a struct with a single boolean field.

| Method                          | Operations/sec | NS/Op    | Bytes/Op | Allocs/Op |
|---------------------------------|----------------|----------|----------|-----------|
| MarshalJSON_bytesBuffer        | 564,442,346    | 2.128 ns/op     | 0 B/op    | 0        |
| MarshalJSON_stringsBuilder      | 553,390,320    | 2.104 ns/op     | 0 B/op    | 0        |
| MarshalJSON_append              | 573,010,333    | 2.137 ns/op     | 0 B/op    | 0        |
| **MarshalJSON_plus**                | **576,073,741**    | **2.102 ns/op**     | 0 B/op    | 0        |
| json.Marshal                    | 9,363,478      | 135.2 ns/op     | 24 B/op   | 1        |
| **ref**                             | **8,564,742**      | **131.7 ns/op**     | 24 B/op   | 1        |

**Output:**  
`{"boolfield":true}`

---

### Benchmark: Pointer Bool Array (Type3)
Performance for a struct with three pointer boolean array fields.

| Method                          | Operations/sec | NS/Op    | Bytes/Op | Allocs/Op |
|---------------------------------|----------------|----------|----------|-----------|
| MarshalJSON_bytesBuffer         | 7,121,541      | 172.5 ns/op     | 192 B/op  | 2        |
| MarshalJSON_stringsBuilder      | 5,653,890      | 225.0 ns/op     | 336 B/op  | 4        |
| **MarshalJSON_append**              | **8,533,594**      | **145.0 ns/op**     | 224 B/op  | 3        |
| MarshalJSON_plus                | 3,349,512      | 372.4 ns/op     | 560 B/op  | 7        |
| json.Marshal                    | 2,727,886      | 481.6 ns/op     | 192 B/op  | 2        |
| **ref**                            | **2,782,668**      | **463.4 ns/op**     | 192 B/op  | 2        |

**Output:**  
`{"pointerboolfield1":[true,null,true],"pointerboolfield2":null,"pointerboolfield3":[null,true,null]}`

---

### Benchmark: Pointer Bool Array (Single Field)
Performance for a struct with a single pointer boolean array field.

| Method                          | Operations/sec | NS/Op    | Bytes/Op | Allocs/Op |
|---------------------------------|----------------|----------|----------|-----------|
| **MarshalJSON_bytesBuffer**         | **16,870,214**     | **71.97 ns/op**     | 64 B/op   | 1        |
| MarshalJSON_stringsBuilder      | 8,640,126      | 132.5 ns/op     | 144 B/op  | 3        |
| MarshalJSON_append              | 16,015,058     | 77.04 ns/op     | 96 B/op   | 2        |
| MarshalJSON_plus                | 7,993,498      | 180.8 ns/op     | 128 B/op  | 3        |
| json.Marshal                    | 4,258,221      | 275.6 ns/op     | 72 B/op   | 2        |
| **ref**                             | **3,661,494**      | **279.6 ns/op**     | 72 B/op   | 2        |

**Output:**  
`{"pointerboolfield":[true,null,true]}`

---

### Benchmark: Pointer Bool (Type3)
Performance for a struct with three pointer boolean fields.

| Method                          | Operations/sec | NS/Op    | Bytes/Op | Allocs/Op |
|---------------------------------|----------------|----------|----------|-----------|
| MarshalJSON_bytesBuffer         | 9,654,776      | 136.9 ns/op     | 192 B/op  | 2        |
| MarshalJSON_stringsBuilder      | 6,872,143      | 190.4 ns/op     | 304 B/op  | 4        |
| **MarshalJSON_append**              | **9,875,238**      | **126.5 ns/op**     | 224 B/op  | 3        |
| MarshalJSON_plus                | 8,467,968      | 147.6 ns/op     | 224 B/op  | 3        |
| json.Marshal                    | 4,227,553      | 291.5 ns/op     | 104 B/op  | 2        |
| **ref**                             | **4,136,919**      | **277.2 ns/op**     | 104 B/op  | 2        |

**Output:**  
`{"pointerboolfield1":true,"pointerboolfield2":null,"pointerboolfield3":true}`

---

### Benchmark: Pointer Bool (Single Field)
Performance for a struct with a single pointer boolean field.

| Method                          | Operations/sec | NS/Op    | Bytes/Op | Allocs/Op |
|---------------------------------|----------------|----------|----------|-----------|
| MarshalJSON_bytesBuffer         | 22,115,852     | 46.32 ns/op     | 64 B/op   | 1        |
| MarshalJSON_stringsBuilder      | 17,097,026     | 75.81 ns/op     | 64 B/op   | 2        |
| **MarshalJSON_append**              | **41,371,674**     | **28.78 ns/op**     | 32 B/op   | 1        |
| MarshalJSON_plus                | 32,664,024     | 33.98 ns/op     | 32 B/op   | 1        |
| json.Marshal                    | 8,083,610      | 148.4 ns/op     | 32 B/op   | 1        |
| **ref**                             | **7,974,753**      | **151.5 ns/op**     | 32 B/op   | 1        |

**Output:**  
`{"pointerboolfield":true}`

---

### Summary

- **Custom Implementations:**  
  At least one custom method outperforms standard `json.Marshal` in every test group.  
- **Memory & Allocations:**  
  Methods using `append` and `bytesBuffer` show lower memory usage and fewer allocations.
- **Method Choice:**  
  Choose the implementation based on your performance and memory usage requirements.


### Mixed Types
Performance for a struct with multiple types (bool, float64, int8, int, string):
| Method | Operations/sec | NS/Op | Bytes/Op | Allocs/Op |
|--------|---------------|--------|-----------|------------|
| bytesBuffer | 3,009,912 | 414.8 | 219 | 4 |
| plus | 1,539,234 | 790.4 | 795 | 12 |
| append | 3,418,110 | 334.5 | 195 | 5 |
| json.Marshal | 2,268,294 | 504.4 | 144 | 2 |

Output: `{"boolfield":true,"float64field":12.34,"int8field":12,"intfield":123,"stringfield":"hello"}`