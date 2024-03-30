package main

import "fmt"

type pavan struct {
	pa int
	ga bool
}

func main() {
	raja := pavan{
		2, true}
	jaji := pavan{
		2, true}
	fmt.Println(jaji)
	fmt.Println(raja)
	var paslice []pavan
	fmt.Println(paslice)
	for i, item := range paslice {
		fmt.Println(i, item)
	}

}
