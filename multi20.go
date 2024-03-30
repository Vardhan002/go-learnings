package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			z := i * j
			fmt.Printf("%d * %d =%d\n", i, j, z)
		}
		fmt.Println()
	}
}
