package main

import "fmt"

func main() {
	var T int
	var x int
	var y int
	fmt.Scanln(&T)
	for i := 1; i <= T; i++ {
		fmt.Scanln(&x)
		fmt.Scanln(&y)
		if x == y {
			println("both are same:")
		} else {
			if x < y {
				fmt.Println(x)
			} else {
				fmt.Println(y)
			}
		}

	}
}
