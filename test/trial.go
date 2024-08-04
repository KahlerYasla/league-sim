package main

import "fmt"

func main() {
	var x int = 42
	var p *int = &x // p now holds the memory address of x

	fmt.Println(x)  // Output: 42
	fmt.Println(*p) // Output: 42 (dereferencing p to get the value)
	fmt.Println(p)  // Output: 0xc00000c028 (memory address)
	fmt.Println(&x) // Output: 0xc00000c028 (memory address)

	*p = 21 // dereferencing p to change the value of x

	fmt.Println()
	fmt.Println(x)  // Output: 21
	fmt.Println(*p) // Output: 21
	fmt.Println(p)  // Output: 0xc00000c028
	fmt.Println(&x) // Output: 0xc00000c028
}
