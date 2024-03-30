package main

import "fmt"

func main() {
	m := map[string]int{
		"James Bond":      32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James Bond"])

	v, ok := m["chocolate"]
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("This is if Print", v)
	}

}
