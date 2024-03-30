package main

import "fmt"

//nth fibonacci number function
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println(fibonacci(7))
}

/*procedure:

fib(0)=0
fib(1)=1
fib(2)=1+0 = 1
fib(3)=1+1 = 2
fib(4)=2+1 = 3
:
:
fib(n)=fib(n-1)+fib(n-2)
*/
