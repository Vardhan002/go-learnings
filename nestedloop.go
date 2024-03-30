package main

import "fmt"

func main() {
	for i := 0; i <= 8; i++ {
		fmt.Printf("The outer loop %d\n", i)
		for j := 0; j < 4; j++ {
			fmt.Printf("\t\t The inner loop %d\n", j)
		}
	}
}
