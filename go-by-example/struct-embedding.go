// Go supports embedding of structs and interfaces to express
// a more seamless compostion of types

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embedding looks like a field without a name
type container struct {
	base
	str string
}

func main() {
	// when creating structs with literals, we have to initalize the embedding explicitly.
	co := container {
		base: base {
			num: 1,
		},
		str: "some name",
	}

	// We can acess the base's fields directly on co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num: ", co.base.num)

	fmt.Println("describe: ", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer: ", d.describe())
}