package main

import "fmt"

func Calculate(i, j *int) {
	if i == nil || j == nil {
		return // handle error and return. Can't operate
	}

	fmt.Println(*i) // read i through the pointer
	*i = 21         // set value directly to the memory i
	fmt.Println(*i) // see the new value of i

	fmt.Println(*j)
	*j = *j / 37    // divide j through the pointer
	fmt.Println(*j) // set value directly to the memory j
}

func main() {
	i := 42
	j := 2701
	Calculate(&i, &j)

	fmt.Printf("After caculate i is: %v\n", i) // you won't get 42
	fmt.Printf("After caculate j is: %v\n", j) // you won't get 2701
}
