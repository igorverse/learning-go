// Go support pointers, allowing you to pass references to values and records within your program

package main

import "fmt"

// Here the argument will be passed by value
func zeroval(ival int) {
	ival = 0
}

// Here the argument will be an int pointer, representing by an *int parameter
// The pointer will be dereferenced to the current value at that address.
func zeroptr(iptr *int) {
	*iptr = 42 // Assigning a value to a dereferenced pointer changes the value at the referenced address
}

func main() {
	i := 1
	fmt.Println("initial i: ", i) // 1
	
	zeroval(i)
	fmt.Println("zeroval: ", i) // 1

	// & operator gives the memory address of i. In other words, it's a pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr: ", i) // 42

	fmt.Println("final i: ", i)

	fmt.Println("pointer: ", &i)

}