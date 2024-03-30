package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%c\t%#U\t%d\n", i, i, i)
	}
}
