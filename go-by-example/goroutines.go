// A goroutine is a lightweight thread of execution

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// caliing the function in the usual way, running it synchronously
	f("direct")

	// To invoke this function in a go routine, use go f(s).
	// This new goroutine will execute concurrently with the
	// calling one
	go f("goroutine")

	// It is possible to start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Now, the two function calls are running asynchronously in
	// separate goroutines.
	time.Sleep(time.Second)
	fmt.Println("done")
}