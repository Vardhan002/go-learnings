package main

import "fmt"

func main() {
	s := "fish"
	v := 0
	c := ""
	for s == c {
		if c == `a` || c == `e` || c == `i` || c == `o` || c == `u` {
			v = v + 1
		}
	}
	fmt.Println("number of vowels=", s)
}
