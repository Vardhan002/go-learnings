package main

import "fmt"

func main() {
	a := 4
	fmt.Println("a =", a)
	fmt.Println("&a =", &a)
	fmt.Printf("a =%T\n", a)
	fmt.Printf("&a =%T\n", &a)
	b := &a
	fmt.Println("b =", b)
	fmt.Printf("&b =%T\n", &b)
	fmt.Println("*b =", *b)
	fmt.Printf("*b =%T\n", *b)
}
