package main

import "fmt"

func main() {
	k := 35
	a := 0
	b := 10
	c := 20
	d := 30
	e := 40
	f := 50
	if a < k {
		fmt.Printf(" %d <= 10\t%d <= 10\n", a, k)
	}
	if b < k {
		fmt.Printf("%d <= 20\t%d <= 20\n", b, k)
	}
	if c < k {
		fmt.Printf("%d <= 30\t%d <= 30\n", c, k)
	}
	if d < k {
		fmt.Printf("%d <= 40\t%d <= 40\n", d, k)
	}
	if e < k {
		fmt.Printf("%d <= 50\t%d <= 50\n", e, k)
	}
	if f < k {
		fmt.Printf("%d <= 60\t%d <= 60\n", f, k)
	}
	fmt.Printf(`input number is "%d"`, k)
}
