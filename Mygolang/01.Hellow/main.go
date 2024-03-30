package main

import "fmt"

var z int

type pavan int

var a pavan

func main() {
	a = 43
	fmt.Println(a)
	fmt.Printf("%T", a)
}
