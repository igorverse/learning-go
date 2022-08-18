// Go supports anonymous function, which can form closures.

package main

import "fmt"

// The function intSeq() returns another function, which is defined
// anonymously in the body of intSeq. The returned function closes
// over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
}