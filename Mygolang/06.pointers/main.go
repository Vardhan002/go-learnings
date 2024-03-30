package main

import "fmt"

func main() {

	fmt.Println("pointers learning")

	var pavan int = 2

	fmt.Println(pavan)

	var ptr1 *int
	fmt.Println(ptr1)

	var ptr = &pavan

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr + 2
	fmt.Println("new value:", *ptr)
	fmt.Println("new one", pavan)

}
