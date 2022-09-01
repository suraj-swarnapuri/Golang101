package main

import (
	"fmt"
	"time"
)

func say(s string){
	for i:= 0; i<5; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	fmt.Println("======Go Routines======")
	// goroutine is a lightweight thread managed by the Go routine
	go say("Hello")
	say("world")

	fmt.Println("=======Channels========")

	chnl := make(chan string)
	go func(){
		for i :=0; i<2; i++{
			chnl <- "a"
			time.Sleep(10 * time.Second)
		}
	}()

	for i:=0; i<2; i++{
		fmt.Println(<-chnl)
	}	
	
	
	fmt.Println("=====Buffered Channel======")
	// a sender can close a channel to indicate that no more values will be sent.
	fib := func(n int, c chan int){
		x,y :=0 ,1
		for i :=0; i<n; i++{
			c <- x
			x,y = y, x+y
		}
		close(c)
	}

	c := make(chan int, 10) // buffer channel
	go fib(cap(c),c)
	for i:= range c{ // recieves values from the channel until closed
		fmt.Println(i)
	}

	fmt.Println("========Select=======")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for{
		select{ // will block until a case is satisfied
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Print(" ")
			time.Sleep(50 * time.Millisecond)
			
		}
	}

}