/*
Q:Write a Go program to swap variables in a list?
ans:Consider we have num1=2, num2=3. To swap these two numbers, we can just write: num1,num2 = num2, num1

The same logic can be extended to a list of variables as shown below:
*/
package main

import "fmt"

func swapContents(listObj []int) {
	for i, j := 0, len(listObj)-1; i < j; i, j = i+1, j-1 {
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}
}
func main() {
	listObj := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(listObj)
	swapContents(listObj)
	fmt.Println(listObj)
}
