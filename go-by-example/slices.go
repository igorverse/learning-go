// Slices is a very powerful data type to work with sequences in Go
// It is not size immutable as the array

package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	// slice support append, wich returns a slice containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// slices can be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	// It can use *slice* operator -> slice[low:high]
	l := s[2:5]
	fmt.Println("sl1: ", l)

	// Get all elements excluding s[5]
	l = s[:5]
	fmt.Println("sl2: ", l)

	// Get all elements starting (and including) from s[2]
	l = s[2:]
	fmt.Println("sl3: ", l)

	// Initialized slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// Slices can be multi-dimensional
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}
