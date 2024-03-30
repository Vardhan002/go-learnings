package main

import "fmt"

func main() {
	vfunc(11, 22, 33, 44, 55)
}

func vfunc(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	divid := 11

	for i, v := range x {
		divid = divid + divid
		divid += v
		fmt.Println(i, v, divid)
	}

	fmt.Println(divid)

}
