package main

import "fmt"

type animals struct {
	name string
	legs int
}
type pet struct {
	animals
	domestic bool
	wild     bool
}

func main() {
	a := pet{
		animals: animals{},
	}
	b := pet{
		animals: animals{
			name: "Dog",
			legs: 4,
		},
		domestic: true,
		wild:     false,
	}
	c := pet{
		animals: animals{
			name: "Tiger",
			legs: 4,
		},
		domestic: false,
		wild:     true,
	}
	d := pet{
		animals: animals{
			name: "Wild Dogs",
			legs: 4,
		},
		domestic: true,
		wild:     true,
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(b.animals, b.name, b.legs, b.domestic, b.wild)
	fmt.Println(d.animals, d.wild)
}
