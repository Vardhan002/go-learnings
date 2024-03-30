package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var number string
	fmt.Fscan(r, &number)
	if len(number) <= 3 {
		fmt.Fprintln(w, len(number))
	} else {
		fmt.Fprintln(w, "More than 3 digits")
	}
}
