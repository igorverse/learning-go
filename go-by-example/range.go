// range iterates over elements in a variety of data structures

package main

import "fmt"

func main() {

	// range to sum the numbers in a slice. It's the same for arrays...
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// range on arrays and slices provides both the index and value for each entry
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// range on map iterates over key-value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
