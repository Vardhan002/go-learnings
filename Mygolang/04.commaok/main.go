package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("pavan is a good buy")

	pavan, err := reader.ReadString()
	if err != nil {
		fmt.Println("this is the error", err.Error())
	}
	fmt.Println("How you are telling ", pavan)
}
