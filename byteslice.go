package main

import (
	"bytes"
	"fmt"
)

func main() {

	sl1 := []byte{'I', 'N', 'T', 'E', 'R', 'V', 'I', 'E', 'W'}
	sl2 := []byte{'I', 'N', 'T', 'E', 'R', 'V', 'I', 'E', 'W'}

	// Use Compare function to compare slices
	res := bytes.Compare(sl1, sl2)

	//this prints equal due to the elements in the slices are equal and same
	if res == 0 {
		fmt.Println("Equal Slices")
	} else {
		fmt.Println("Unequal Slices")
	}
	sl1 = []byte{'I', 'N', 'T', 'E', 'R', 'V', 'I', 'E', 'W'}
	sl2 = []byte{'B', 'E', 'T', 't', 'E', 'R', 'V', 'I', 'E', 'W'}

	ser := bytes.Compare(sl1, sl2)

	//this prints unequal due to the elements in the slices are not equal and same
	if ser == 0 {
		fmt.Println("Equal Slices")
	} else {
		fmt.Println("Unequal Slices")
	}

}
