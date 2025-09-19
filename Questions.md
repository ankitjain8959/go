# Key Differences — Go vs Java
Go is preferred for low-latency, lightweight Lambda functions, while
Java is preferred for complex, enterprise-grade logic with rich libraries.

Go is preferred when you need very low cold start times.


# What is `cold start` vs `warm start`?
A `cold start` occurs when a serverless function (AWS Lambda) is invoked for the first time or after a period of inactivity. <br>
The cloud provider must allocate resources, which can lead to increased latency for that initial request. <br>

It involves: <br>
- Downloading the code
- Setting up the environment

Please note, that Cloud provider (AWS) will not charge for the cold start time, but it does add latency to your overall invocation duration.
After the invocation completes, the execution environment is frozen. To improve resource management and performance, Lambda retains the execution environment for a period of time. During this time, if another request arrives for the same function, Lambda can reuse the environment. This second request typically finishes more quickly, since the execution environment is already fully set up. This is called a `warm start`.

Go Cold Start time is typically lower than Java because:
- Go binaries are compiled to machine code, which allows for faster startup times.

Go cold start time is typically around 100-200 milliseconds, while Java can take 500 milliseconds or more.

# Where Are Go Dependencies Stored on the Local Machine?
- Go automatically downloads and caches modules the first time you run: `go build` / `go get` / `go mod tidy`
- Go caches all module dependencies in your local system under `$GOPATH/pkg/mod` like, Maven uses `~/.m2/repository` for Java.

# What are the dependency config files used in Go?
- `go.mod`: This file defines the module's name and its dependencies. It specifies the required versions of each dependency.
- `go.sum`: This file contains checksums (hash) for the module's dependencies, ensuring that the correct versions are used and preventing tampering.

Like, Maven uses `pom.xml` for Java.

# What is Pointers? What is referencing & dereferencing? Why do we need pointers?
- A pointer is a variable that stores the memory address of another variable.

| Term                    | Meaning                                                                  |
| ----------------------- | ------------------------------------------------------------------------ |
| **Referencing** (`&`)   | Getting the memory **address** of a variable (i.e., creating a pointer). |
| **Dereferencing** (`*`) | Accessing or changing the **value at the pointer’s address**.            |

```
package main

import "fmt"

func main() {
    x := 10
    ptr := &x           // referencing: get address of x
    fmt.Println(ptr)    // prints something like 0xc000012090 (the memory address)
    fmt.Println(*ptr)   // dereferencing: prints 10 (value at that address)

    *ptr = 20           // changes the value at the address
    fmt.Println(x)      // prints 20 — x was updated
}
```
Pointers allow you to: <br>
- Modify a variable inside a function (because Go passes everything by value)
- To avoid copying big data (e.g., large structs).

# Go: Pass by Value or Pass by Reference?
| Term                  | What it Really Means                                                                                                           |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| **Pass by Reference**     | The function receives a **reference (pointer)** to the variable. Changing it inside the function **does affect the original**. |
| **Pass by Value** | The function receives a **copy** of the variable. Changing it inside the function **won’t affect the original**. | 

> Go is always Pass by Value. <br>

```
func UnderstandPointers() {

	var x int = 10
	var p *int = &x // p is a pointer to x, it holds the address of x

	fmt.Println("Value of x:", x)                    // Output: Value of x: 10
	fmt.Println("Address of x:", &x)                 // Output: Address of x: 0x... (some memory address)
	fmt.Println("Value of p:", p)                    // Output: Value of p: 0x... (address of x)
	fmt.Println("Value at address p points to:", *p) // Output: Value at address p points to: 10

	passByValue(p)                                    // Here, we are passing the pointer p to the function, which allows us to modify the value of x indirectly. but the value of p remains unchanged.
	fmt.Println("Value of p:", p)                     // Output: Value of p: 0x... (address of x)
	fmt.Println("Value of x after function call:", x) // Output: Value of x after function call: 20
}

func passByValue(y *int) {
	fmt.Println("Value of y inside function:", y)
	*y = 20
}
```
Note: <br>
- When you pass a pointer to a function, you're passing a copy of the pointer (still pass-by-value) & the value of the pointer ( i.e. memory address) will remain unchanged.
- But that pointer points to the original value, so the function can modify the original data using dereferencing.