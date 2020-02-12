package main

import (
	"fmt"
	"math"
)

const s string = "constant" // A const statement can appear anywhere a var statement can.

func main() {
	fmt.Println(s)

	const n = 500000000 // A numeric constant has no type until it’s given one, such as by an explicit conversion.

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
