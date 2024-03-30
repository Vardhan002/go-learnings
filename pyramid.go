package main

import "fmt"

func main() {

	rows := 10
	space := rows - 1

	for j := 1; j <= rows; j++ {
		for i := 1; i <= space; i++ {
			fmt.Print(" ")
		}
		space--
		for i := 1; i <= 2*j-1; i++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}
