// Go's strcuts are typed collections of fields. They're useful
// for grouping data together to form records

package main

import "fmt"

type person struct {
	name string
	age int
}

// it constructs a new person struct with the given name
func newPerson(name string) *person {

	// You can safely return a pointer to local variable as a local variable will survive
	// the scope of the function
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	// it creates a new struct
	fmt.Println(person{"Bob", 20})

	// you can create a struct with the fields name
	fmt.Println(person{name: "Alice", age: 30})
	// Ommited fields will be zero-valued
	fmt.Println(person{name: "Heitor"})
	
	// & yields a pointer to the struct
	fmt.Println(&person{"Igor", 23})

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // access fields with a dot

	sp := &s // the pointers are automatically dereferenced
	fmt.Println(sp.age) // So, you can use dots with struct pointers

	sp.age = 51 // Structs are mutalbe
	fmt.Println(sp.age)



}