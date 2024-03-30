package main

import "fmt"

func main() {
	raja("pavan")
	g := dude("myage:")
	fmt.Println(g)
	t := dft(22)
	fmt.Println(t)
}

func raja(s string) {

	fmt.Println(s)
}

func dude(s string) string {
	return fmt.Sprint(s, 22)

}
func dft(s int) int {
	a := s
	return a

}
