package main

import "fmt"

const (
	a = 2022 + iota
	b = 2022 + iota
	c = iota
	d = iota
)

func main() {
	fmt.Println("\n", a, "\n", b, "\n", c, "\n", d, "\n")
}
