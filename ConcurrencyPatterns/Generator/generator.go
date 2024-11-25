package main

import "fmt"
//import "time"
import "sync"

var wg sync.WaitGroup
func main() {
	wg.Add(1)
	c := generator()
	wg.Add(1)
	go consumer(c)

	wg.Wait()
	//time.Sleep(1*time.Second)
}

func generator() chan int {
	defer wg.Done()
	out := make(chan int)

	go func() {
		fmt.Println("I am generator")
		for i:=0; i<20; i++ {
			out <- 2*i
		}
		close(out)
	}()

	return out
}

func consumer(c <-chan int) {
	defer wg.Done()
	fmt.Println("I am a consumer")
	for {
		select {
		case v, ok := <-c:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}

}