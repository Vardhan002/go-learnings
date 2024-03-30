package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("anonymous function")

	}() //parens must used to return paramerets that are passed into arguments
	//if () are missing then"./prog.go:7:2: func() {…} (value of type func()) is not used"
	func(x int) {
		fmt.Println("i can print", x)
	}(47)
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
}
func foo() {
	fmt.Println("i am foo")
}
