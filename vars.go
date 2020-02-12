package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)
	
	var b, c int = 1,2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)

	var e int // Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
	fmt.Println(e)
	
	f := "apple" // The := syntax is shorthand for declaring and initializing a variable
	fmt.Println(f)
}
