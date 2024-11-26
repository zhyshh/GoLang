package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	done := make(chan interface{})
	//defer close(done)

	cows := make(chan interface{}, 100)
	pigs := make(chan interface{}, 100)

	//two producers, produce infinitely
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Pig producer done")
				return
			case pigs <- "oink":
			}
		}
	}()
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Cow producer done")
				return
			case cows <- "moo":
			}
		}
	}()

	wg.Add(1)
	go consumePig(done, pigs)
	wg.Add(1)
	go consumeCow(done, cows)

	time.Sleep(20 * time.Second)
	fmt.Println("Time's up, closing")
	close(done)

	wg.Wait()

}

func consumeCow(done <-chan interface{}, cows <-chan interface{}) {
	defer wg.Done()
	for cow := range orDone(done, cows) {
		time.Sleep(1 * time.Second)
		fmt.Println(cow)
	}
	fmt.Println("Cow consumer done")
}

func consumePig(done <-chan interface{}, pigs <-chan interface{}) {
	defer wg.Done()
	for pig := range orDone(done, pigs) {
		time.Sleep(1 * time.Second)
		fmt.Println(pig)
	}
	fmt.Println("Pig consumer done")
}

func orDone(done, c <-chan interface{}) <-chan interface{} {
	relayStream := make(chan interface{})
	go func() {
		defer close(relayStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case relayStream <- v:
				case <-done:
					return //to prevent infinite block of above case relayStream<-v:
				}
			}
		}
	}()
	return relayStream
}
