// var declares 1 or more variables
// Go can infer the type of initialized variables
// := is a shorthand to declare and initialize a variable: var name = 'igor' -> name := 'igor'
// variables without a initialization are zero-valued

package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
