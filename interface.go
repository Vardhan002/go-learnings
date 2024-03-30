package main

import (
	"fmt"
)

type peson struct {
	first string
	last  string
	age   int
}

type place struct {
	peson
	staying string
	ltk     bool
}

func (s place) speak() {
	fmt.Println("I am", s.first, s.last, "staying in -", s.staying)
}

func (p peson) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
}

func hello(h human) {
	switch h.(type) {
	case peson:
		fmt.Println("I was passed into hello", h.(peson).first)
	case place:
		fmt.Println("I was passed into hello", h.(place).first)
	}
	fmt.Println("I was passed into hello", h)
}

type house int

func main() {
	sa1 := place{
		peson: peson{
			"vivek",
			"vardhan",
			24,
		},
		staying: "kurnool",
		ltk:     true,
	}

	sa2 := place{
		peson: peson{
			"siva",
			"ram",
			38,
		},
		staying: "kurnool",
		ltk:     true,
	}

	p1 := peson{
		first: "abhi",
		last:  "ram",
		age:   6,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	hello(sa1)
	hello(sa2)
	hello(p1)

	// conversion
	var x house = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
