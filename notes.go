package main

import "fmt"

func main() {
	notes := [10]int{2000, 500, 200, 100, 50, 20, 10, 5, 2, 1}
	var amount int
	fmt.Print("Enter the Total Amount of Cash = ")
	fmt.Scanln(&amount)

	temp := amount
	for i := 0; i < 10; i++ {
		fmt.Println(notes[i], " Notes = ", temp/notes[i])
		temp = temp % notes[i]
	}
}
