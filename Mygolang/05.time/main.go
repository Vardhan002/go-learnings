package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("using time")
	mytime := time.Now()
	fmt.Println(mytime)

	fmt.Println(mytime.Format("monday"))
}
