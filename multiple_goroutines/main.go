package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs chan int, result chan int) {
	for j := range jobs {
		fmt.Println("job id: ", id, "-", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}

func main() {
	maxjobs := 5
	jobs := make(chan int, maxjobs)
	result := make(chan int, maxjobs)

	for i := 1; i <= 3; i++ {
		go workers(i, jobs, result)
	}

	for i := 0; i < maxjobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < maxjobs; i++ {
		fmt.Println("Result: ", <-result)
	}
}
