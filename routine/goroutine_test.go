package main

import (
	"fmt"
	"testing"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "end     job", j)
		result <- j * 2
	}
}
func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 0; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 10; j <= 15; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}

}
