package main

import "fmt"

const a = iota
const b = iota
const c = iota
const (
	d = iota + 0.5
	e = iota
	f
	g
)

func main() {
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d, "e=", e, "f=", f, "g=", g)
}
