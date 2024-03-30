package main

import "fmt"

func main() {
	foo()
	bar("james")
	s1 := woo("MoneyPenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}
func foo() {
	fmt.Println("Hello from foo")
}
func bar(S string) {
	fmt.Println("Hello,", S)
}
func woo(St string) string {
	return fmt.Sprint("hello from woo,", St)
}
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `,says "Hello"`)
	b := false
	return a, b
}
