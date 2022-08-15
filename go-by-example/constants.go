// const declares a constant value and can appear anywhere a var statement can
// a numeric constant has no type. It will have if you give one and by an explict conversion

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
