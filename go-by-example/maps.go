// Maps (sometimes called hashes or dicts in other languages) are Go's built-in associative data type

package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Set key/value pairs using typeica name[key] = value syntax
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	// As the name sugests, it removes a key-value pair from a map
	delete(m, "k2")
	fmt.Println("map: ", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	// Initialized map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}
