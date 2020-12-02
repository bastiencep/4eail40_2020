package main

import (
	"fmt"
	"time"
)

// Each worker is a loop awaiting tasks from the job channel.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	const numWorkers = 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Creating the workers.
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Creating the jobs and pushing them into the jobs queue.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// Closing the jobs channel will inform workers that no more job is incoming.
	close(jobs)

	// Fetching the results.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
