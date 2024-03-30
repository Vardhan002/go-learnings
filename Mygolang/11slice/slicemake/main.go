package main

import "fmt"

func main() {
	z := make([]int, 3, 100)
	z[0] = 23
	z[2] = 33
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

}
