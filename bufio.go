package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" //converting the type
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the Year you were born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64) //converting string to integer by using strconv and give base 10 is of decimal value and 64 is size 0f type int64
	fmt.Printf("You will be %d years old at the end of this year", 2023-input)

}
