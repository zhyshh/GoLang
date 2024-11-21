//alternative implementation 

package main

import "fmt"

func worker(in chan int) chan int {
	out := make(chan int)

	go func() {
		v := <- in
		out <- v+1
	}()

	return out
}

const N = 100000
func main() {
	c := make(chan int)
	in := c
	var out chan int

	for i:=0; i<N; i++ {
		out = worker(in)
		in = out
	}
	c<-0
	fmt.Println(<-out)
}