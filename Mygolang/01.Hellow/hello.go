package main

import "fmt"

func main() {
	z := [2]int{1}
	fmt.Println(z)

	v := []int{1, 2, 3, 4}
	fmt.Println(len(v))
	fmt.Println(cap(v))

	a := append(v, 6, 7)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := append(a, 8)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	c := append(b, 5)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	d := append(c, 9)
	fmt.Println(len(d))
	fmt.Println(cap(d))

}
