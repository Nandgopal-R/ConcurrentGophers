package main

import (
	"fmt"
	"time"
)

func printMessage(msg chan string) {
	time.Sleep(time.Second * 2)
	msg <- "Hello from printMessage function"
}

func main() {
	msg := make(chan string)
	go printMessage(msg)
	fmt.Println("Hello from main")
	fmt.Println(<-msg)
}
