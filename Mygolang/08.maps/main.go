package main

import "fmt"

func main() {
	fmt.Println("learning maps")
	m := []int{1, 2, 3}
	pmaps := map[string]string{
		"pa": "pavan",
		"aj": "ajesh",
	}

	// fmt.Println("pavan short name: ", pmaps)
	// fmt.Println("pavan short name: ", pmaps["pa"])

	// if p, ok := pmaps["aj"]; ok {
	// 	fmt.Println(p, ok)
	// }

	pmaps["bh"] = "bsuresh"
	//fmt.Println(pmaps["bh"])
	// for _, p := range pmaps {
	// 	fmt.Println(p)
		
	// 	//fmt.Println(ok)
	// }

	// for i, P := range p {
	// 	fmt.Println(i, z)
	// }

}
