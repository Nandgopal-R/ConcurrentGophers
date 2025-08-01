package main

import (
	"fmt"
	"time"
)

func greet() {
	fmt.Println("Hello from greet function")
}

func main() {
	go greet()                              // This creates a go routine that makes function greet run concurrently with the rest of main function
	fmt.Println("Hello from main function") // prints this from main goroutine
	time.Sleep(time.Second)                 // main goroutine sleeeps for 1 sec to give time for greet()
}
