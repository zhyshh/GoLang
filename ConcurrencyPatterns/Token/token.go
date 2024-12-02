package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a buffered channel with capacity 2, representing the number of tokens
	tokens := make(chan struct{}, 2)

	// Fill the channel with tokens
	for i := 0; i < cap(tokens); i++ {
		tokens <- struct{}{}
	}

	var wg sync.WaitGroup

	// Launch multiple goroutines to perform work
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Acquire a token
			<-tokens
			fmt.Printf("Goroutine %d acquired a token\n", id)

			// Simulate work
			time.Sleep(time.Second)

			// Release the token
			fmt.Printf("Goroutine %d released a token\n", id)
			tokens <- struct{}{}
			
		}(i)
	}

	wg.Wait()
}

/*
Explanation:
 Buffered Channel:
    The tokens channel acts as a semaphore, limiting the number of concurrent operations. The capacity of the channel determines the maximum number of tokens available.
 Acquiring a Token:
    A goroutine acquires a token by reading from the tokens channel. If no tokens are available, the goroutine will block until one becomes available.
 Releasing a Token:
    After completing its work, the goroutine releases the token by sending a value back to the tokens channel.
 WaitGroup:
    The WaitGroup ensures that all goroutines finish their work before the program exits.
*/
