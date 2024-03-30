package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("happy")
	case false:
		fmt.Println("sad")
	case false:
		fmt.Println("tired")
	}
}
