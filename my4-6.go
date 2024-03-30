package main

import "fmt"

func main() {
	m := make([]string, 0)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	y := []string{"water", "juice", "milk", "tea", "coffee", "milkshake"}
	m = append(y)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	for i, v := range m {
		fmt.Println(i, v)
	}
}
