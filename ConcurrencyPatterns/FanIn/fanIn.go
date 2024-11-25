package main

import "fmt"

func main() {
	c0 := producer("gooooo")
	c1 := producer("booooo")
	c2 := producer("fooooo")
	c3 := producer("mooooo")

	out := consumer(c0,c1,c2,c3)
	
	for v := range out {
		fmt.Println(v)
	}
}

func producer(s string) chan string {
	//defer wg.Done()
	out := make(chan string)

	go func() {
		fmt.Println("I am producer for",s)
		for {
			out <- s
		}
		
	}()

	return out
}

func consumer(c0, c1, c2, c3 <-chan string) chan string{

	out := make(chan string)
	reader := func(in <-chan string){
		for v := range in {
			out<-v
		}
	} 

	go reader(c0)
	go reader(c1)
	go reader(c2)
	go reader(c3)
	return out

}