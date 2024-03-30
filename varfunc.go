package main

import "fmt"

func main() {
	x := sum(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("total is", x)
}
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("index =", i, " ", "value =", v, " ", "Totalsum =", sum)
	}
	return sum
}
