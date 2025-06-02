package basics

import "fmt"

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
