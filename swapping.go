package main

import "fmt"

var x = 5
var y = 9

func main() {
	fmt.Println("x =", x, "y =", y)
	x, y = y, x
	fmt.Printf("x =%v\ny =%v\t", x, y)
}
