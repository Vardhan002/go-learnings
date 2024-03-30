package main

func P(v *int) { print(*v) }

func main() {
	for i := range [3]int{} {
		defer P(&i)
	}
}
