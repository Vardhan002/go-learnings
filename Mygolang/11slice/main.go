package main

import "fmt"

func main() {

	var parra [3]int

	parra[0] = 1
	parra[1] = 2
	parra[2] = 3

	fmt.Println(parra)

	P := []int{1, 2, 3, 4, 6}
	s := len(P)
	//k := []int{1111, 222, 333}
	//z := []int{102, 202, 303}

	//P = append(P)
	fmt.Println(P)
	fmt.Println(s)
	// fmt.Println(len(p))
	// P = append(P, z...)
	// fmt.Println(P)
	// P = append(P[:5], P[7:]...)
	// fmt.Println(P)
	// fmt.Println(P)
	// fmt.Println(P[2])
	// fmt.Println(P[3])
	// fmt.Println(P[3:])
	// P = append(P, 11, 22, 33)
	// fmt.Println(P)

	// for i, k := range P {
	// 	fmt.Println(i, k)

	// 	for i = 0; i <= 3; i++ {
	// 		fmt.Println(i, P[i])
	// 	}
	// }

}
