package main

import (
	"fmt"
)

func main() {
	var x int
	y := [3]int{3, 5, 2}
	x = y[0]
	fmt.Printf("%v%v\t", x, y[1:])
	fmt.Println()
}
