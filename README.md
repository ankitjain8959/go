# Introduction: Go Programming Language
- Go was created at Google in 2007 and was made public in 2009.

# Download & Install Go
https://go.dev/doc/install

# Dependency Management/Tracking in Go: `go.mod` & `go.sum` files
1. `go.mod` file <br>
It is used for dependency management in GO, similar to how `pom.xml` file manages dependencies in a Java-Maven project. <br>
It tells Go:
- What your project is called like <groupId> + <artifactId> in `pom.xml`
- Which version of Go you're using like (<java.version>...</java.version>) in `pom.xml`
- Which external packages + their versions your code needs like (<dependencies>...</dependencies>) in `pom.xml`

Sample `go.mod` file
```
module github.com/yourusername/myapp

go 1.20

require (
    github.com/gin-gonic/gin v1.9.0
    golang.org/x/crypto v0.9.0
)
```

Command to create `go.mod` file
```
go mod init <module>
```

2. `go.sum` file <br>
It is used to ensure the security and integrity of your dependencies in Go. <br>
It stores cryptographic checksums (hashes) for every dependency (module version) your project uses. <br>
It gets created automatically - the first time when you run a commands like: <br>
`go mod tidy` or `go get` or `go build` or  `go test` <br>

Note: `go.sum` is like Maven's `.m2/repository` metadata or checksum info - it ensures downloaded modules haven't been tampered with.

Sample `go.sum` file
```
github.com/gin-gonic/gin v1.9.0 h1:abc123...        # content hash
github.com/gin-gonic/gin v1.9.0/go.mod h1:def456... # The module's go.mod hash
```
There will be 2 entries for each dependency:
- One for the actual code of the module (to make sure that The actual source code of the module is safe & unchanged)
- One for the module's `go.mod` file (to make sure that the go.mod file inside that module is safe & unchanged)

**IMPORTANT**: <br>
You _do not need to manually manage `go.mod` and `go.sum` files_, unless you want to:
- Rename your module (edit `module` line in `go.mod`)
- Change the Go version (edit `go` line in `go.mod`)

It will be managed automatically when you run a commands like: <br>
`go mod tidy` or `go get` or `go build` or  `go test` <br>


# Packages & Modules in GO
`package`
- A package is a collection of related Go source files in the same directory that are compiled together.
- This of it like, a folder that groups related `.go` files under a name.

Example: <br>
//mathutils/math.go
```
package mathutils

func Add(a, b int) int {
    return a + b
}
```
To use this package in another file:
```
import "your-module-name/mathutils"
```

`module`
- A module is a collection of packages, versioned together, with a `go.mod` file at the root.
- Think of it like, a project or repository that may contain many packages.

Example: <br>
```
your-app/           ← Module (with go.mod)
├── go.mod
├── main.go         ← part of package main
└── utils/
    └── math.go     ← part of package utils
```

Inside `go.mod` the module will be defined as:
```
module github.com/yourusername/your-app
```

# Pointers in Go
- A pointer is a special kind of variable that is not only used to store the memory addresses of other variables but also points where the memory is located and provides ways to find out the value stored at that memory location.
- It is generally termed as a special kind of Variable because it is almost declared as a variable but with *(dereferencing operator).
- The type *T is a pointer to a T value. Its zero value is nil.
```
> Normal variable
var x int = 567

> Declaration of pointer p
var p *int

> Initialization of pointer p
p = &x

whereas, <br>
x => 567
p => 0x414020
&x = 0x414020

To access the value stored in the variable which the pointer points to.
* operator is also termed as the `value at the address of`.
*p => 567
```

- `* Operator` also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
- `& Operator` termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

![pointers.png](images\pointers.png)

# Type in Go
- The `type` keyword is used to **define a new data type** or **give a name to an existing type** in Go. <br>

- The `type` keyword is used to:
    > Define your own data types <br>
      Create types based on existing ones (like, int, string, bool, float64 etc) <br>
      Group fields together (using struct) <br>
      Describe behavior (using interface) <br>

```
> Built in type (eg. int)
var age int

> Custom type based on a built-in type
type Age int    // Defining a new type Age based on int
var a Age = 25

> struct type
type Person struct {
    Name string
    Age  int
}
// This defines a new data type Person, made of two fields.
```

# Struct in Go
- A `struct` is a composite data type that groups together variables (fields) under a single name
- A **struct is a collection of fields** and these fields are accessed using a dot
- Variables inside the type if starts with uppercase letter, are visible and accessible outside the package
- Every variable has a **default value** (also known as the **zero value**) if not explicitly initialized

```
> Structure for User type

type User struct {
	FirstName   string    `json:"firstName,omitempty,bson:"name,omitempty"`
	LastName    string    `json:"lastName"`
	PhoneNumber string    `json:"phoneNumber"`
	Age         int       `json:"age,omitempty"` 
	BirthDate   time.Time `json:"birthDate"` // time.Time is a built-in type in Go for handling date and time
}

// The struct tags (e.g., json or bson etc) are used to control how the struct fields are serialized/deserialized to.
	
// Using json tag specify how the field will be named when the struct is converted to JSON. For example, StatusCode will appear as statusCode in the JSON output.
	
// Using bson tag specify how the field will be named when the struct is converted to BSON (Binary JSON, used in MongoDB).
	
// Using omitempty option allows omitting the field from JSON/BSON etc if it has aan empty value (defined as false, 0, nil pointer, nil interface value & any array, slice, map, or string of length zero.)
```

# Interface in Go
- An `interface` is a type that specifies a contract or a set of method signatures (i.e. behaviours)
- If a type implements all the methods in the interface, it satisfies that interface — automatically, without explicitly declaring it.

```
type Speaker interface {
    Speak() string
}

// A struct type
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

func makeItSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    d := Dog{}
    c := Cat{}
    makeItSpeak(d) // Output: Woof!
    makeItSpeak(c) // Output: Meow!
}
```
Key Points:
- Dog and Cat types both implement the Speaker interface by providing a Speak method and therefore both `Dog` and `Cat` are considered as `Speaker`
- No need to write implements Speaker - it's automatic in Go (Implicit interface implementation)

# Data Types & their default values (also known as, Zero values)

| **Type**              | **Default Value** | **Explanation**                                |
|-----------------------|-------------------|------------------------------------------------|
| `int`, `int32`, `int64`, etc. | `0`               | All integer types default to zero              |
| `float32`, `float64`  | `0.0`             | All float types default to 0.0                 |
| `bool`                | `false`           | Default for booleans                           |
| `string`              | `""`              | Empty string                                   |
| `pointer` types       | `nil`             | Any pointer type defaults to nil               |
| `slice`               | `nil`             | Not an empty slice (`[]`), but `nil`           |
| `map`                 | `nil`             | Must be initialized with `make` before use     |
| `channel`             | `nil`             | Must use `make` to create before using         |
| `interface{}`         | `nil`             | Empty interface defaults to nil                |
| `array`               | Elements are zero | Fixed-size array, elements are zero-initialized|
| `struct`              | Fields are zero   | Each field gets its own zero value             |
| `function`            | `nil`             | Functions can be nil too                       |


# Commands in GO
1. Create `go.mod` file
```
go mod init <module>
```

2. Clean & Sync dependencies
```
go mod tidy
```

Note:
- Add missing dependencies to `go.mod` & `go.sum` (if you're importing packages in code but forgot to run go get)
- Remove any unused module dependencies from `go.mod` and `go.sum` (if you've deleted code but leftover entries remain)

3. Remove object files from current module
```
go clean
```
- Removes object files and cached test results from the current module.
- It does not affect the module cache (where dependencies are stored).

4. Delete local cache of downloaded modules
```
go clean -modcache
```
- This command deletes your local Go module (dependencies) cache which is stored in the `$GOPATH/pkg/mod` directory.
- It removes all downloaded dependencies, forcing Go to re-download them the next time you build or tidy. Useful if your cache is corrupted or you want to clear out old dependencies.

5. Run code
```
go run hello.go
```

6. Test code
```
go test

go test -v    // Increase verbosity of test output
```



# Writing Test
- Ending a file's name with `_test.go` tells the `go test` command that this file contains test functions.
- The file ending with `_test.go` will be excluded from regular package builds `go build` but will be included when the `go test` command is run.
- The Go language provides a minimal yet complete package called `testing` https://pkg.go.dev/testing#pkg-overview

<br>

- Test functions must start with `Test`, followed by a suffix whose first word is uppercase (example: TestXxxx) 
- Test functions in Go receive only one parameter, and in this case, it’s a `pointer of type testing.T`
- If the test file is in the same package as the one being tested, it may refer to unexported identifiers within the package <br>
& If the file is in a separate "_test" package, the package being tested must be imported explicitly and only its exported identifiers may be used. This is known as `black box testing`

```
> math_test.go

package math
import "testing"

func TestXxx(t *testing.T) {}
```