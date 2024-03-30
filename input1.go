package main

import "fmt"

var c float64
var f float64

func main() {
	fmt.Print("enter degrees:")
	fmt.Scanln(&f)
	c = (f - 32) * 5 / 9
	fmt.Println("conversion f to c", c)
}
