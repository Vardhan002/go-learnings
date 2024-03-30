package main

import (
	"fmt"
)

func main() {
	var t, a, b int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		//for j := i; j <= i; j++ {
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)
		fmt.Println("a =", a)
		fmt.Println("b =", b)
		fmt.Println(a + b)

		//}
	}

}
