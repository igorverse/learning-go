// Go supports recursive functions. Here's a classic example.

package main

import "fmt"

// 0 is the base case and n * fact(n - 1) is the recursive case
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n - 1)
}

func memoizedFactorial(n int) int {
	return factHelper(n, map[int]int{ 0: 1 })
}

func factHelper(n int, memoization map[int]int) int {
	if value, found := memoization[n]; found {
		return value
	}

	memoization[n] = n * factHelper(n - 1, memoization)

	return memoization[n]
}

func main() {
	fmt.Println(fact(7)) // 5040
	fmt.Println(memoizedFactorial(7)) // 5040

	// Clousures can also be recursive, but this requires the closure 
	// to be declared with a typed var explicitly before it's defined
	var fib func(n int) int

	// fib was previously declared in main, so go knows which function to call with fib here
	fib = func (n int) int {
		if n == 0 {
			return 0
		}

		if n == 1 {
			return 1
		}

		return fib(n - 1) + fib(n - 2)
	}

	fmt.Println(fib(7)) // 13
}