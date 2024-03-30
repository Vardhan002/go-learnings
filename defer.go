package main

import "fmt"

func main() {
	defer foo()
	bar()
}
func foo() {
	fmt.Println("i was to execute first")
}
func bar() {
	fmt.Println("i will be executed last")
}
