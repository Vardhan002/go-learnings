package main

import "fmt"

func main() {
	am := func() {
		fmt.Println("good morning")
	}
	am()
	pm := func(x int) {
		fmt.Println("good afternoon it above:", x, "pm")
	}
	pm(12)
}
