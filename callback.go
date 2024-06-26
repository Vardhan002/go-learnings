package main

import (
	"fmt"
)

func main() {
	t := evenSum(sum, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(t)
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
		fmt.Println("n=", n)

	}
	return n
}

func evenSum(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
			fmt.Println("xi=", xi)
		}
	}
	total := f(xi...)
	fmt.Println("total=", total)

	return total

}
