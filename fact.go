/*
Q:Write a GO Program to find factorial of a given number.
ans:Factorial of a number is the product of multiplication of a number n with every preceding number till it reaches 1. Factorial of 0 is 1.
Example:
fact(1) = 1
fact(3) = 3 * 2 * 1 = 6
fact(5) = 5 * 4 * 3 * 2 * 1 = 120
*/

package main

import "fmt"

//factorial function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(7))
}
