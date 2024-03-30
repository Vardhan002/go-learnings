package main

import (
	"fmt"
	"time"
)

func main() {
	go sampleRoutine()
	fmt.Println("Started Main")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Main")
}

func sampleRoutine() {
	fmt.Println("Inside Sample Goroutine")
}

/*
In this code, we see that the sampleRoutine() function is called by specifying the keyword go before it.
When a function is called a goroutine, the call will be returned immediately to the next line of the program statement
which is why “Started Main” would be printed first and the goroutine will be scheduled and run concurrently in the background.
The sleep statement ensures that the goroutine is scheduled before the completion of the main goroutine.
*/
