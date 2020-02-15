package main
import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 { // You can have an if statement without an else.
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { // A statement can precede conditional; any variables declared in thi statemnet are available in all branches.
		fmt.Println(num, "is negative")
	} else if num < 10 { 
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
