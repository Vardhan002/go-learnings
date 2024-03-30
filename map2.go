package main

import "fmt"

func main() {
	m := map[string]int{
		"hello":            1,
		"hi":               2,
		"welcome":          3,
		"have a great day": 4,
	}
	fmt.Println(m)
	fmt.Println(m["welcome"])
	m["good morning"] = 5
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("%v\t%v\t\t", k, v)

	}

}
