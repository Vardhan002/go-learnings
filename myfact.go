package main

import "fmt"

func main() {
	for n := 5; n >= 1; n-- {
		j := n*n - 1
		total := j * n
		fmt.Println(total)
	}
}

//chages must be made
