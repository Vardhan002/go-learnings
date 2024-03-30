// Go program to demonstrate the example
// of variable shadowing
package main

import (
	"fmt"
)

func main() {
	x := 0

	// Print the value
	fmt.Println("Before the decision block, x:", x)

	if true {
		x := 1
		x++
	}

	// Print the value
	fmt.Println("After the decision block, x:", x)
	//after resolving
	if true {
		x = 1
		x++
	}

	// Print the value
	fmt.Println("after removing : colon in if condition")
	fmt.Println("After the decision block, x:", x)
}
