package main

import "fmt"

func main() {
	x := 33
	for {
		x++
		if x > 122 {
			break
		}
		fmt.Printf("%v\t%c\n", x, x)
	}
}
