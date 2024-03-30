package main

import "fmt"

func main() {
	m := map[string][]string{
		"pavan": []string{`pa1`, `ganesh`},
		"raja":  []string{`sajj`, `sure`},
	}
	//fmt.Println(m["pavan"])

	for k, l := range m {
		fmt.Println("value of", k, l)
		// for i, l2 := range l {
		// 	fmt.Println(i, l2)
		// }
	}
}
