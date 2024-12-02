package main
import (
	"fmt"
	"time"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)

		time.Sleep(time.Second)// simulate time-consuming task
		results <- j * j// square the number and send the result back

		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	// Setup the queue
	numJobs := 10
	numWorkers := 3
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	// Start the workers.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs to the workers and close the jobs channel when done.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)


	// Wait for all the workers to finish.
	go func() {
		wg.Wait()
		close(results)
	}()


	// Collect and print the results.
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

}
