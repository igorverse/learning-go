// Go has built-in support for multipe return values.
// This is used often, for example, to return both
// result and error values from a function

package main

import "fmt"

// (int, int) signature shows that the function returns
// 2 ints
func values() (int, int) {
	return 3, 7
}

func main() {

	// Using multiple assignment to return 2 different values
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	// Using blank identifier (_) to return 
	// a subset of the returned values
	c, _ := values()
	fmt.Println(c)
}