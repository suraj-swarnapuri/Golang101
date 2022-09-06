package main

import (
	"fmt"
	"time"
)

// example https://github.com/github/ghae-quota-monitor/blob/31e659553fb4ce39c6c773e79ea69d617cd94c76/internal/config/config.go#L110

func worker(id int) {
    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main(){

	// Go routines are lightweight threads of execution that are managed by the go runtime
	fmt.Println("===Go Routines===")

	/*go func(){
		for i:=0; i<4; i++{
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Hello from a go routine")
		}
	}()

	for i:=0; i<4; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello from main thread")
	} */

	// sync go routines
	/*
	var wg sync.WaitGroup
	for i:=0; i<4; i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			worker(i)
		}(i)
	}
	wg.Wait()*/



/*
	a := make(chan bool, 10) // buffered channel
	b := make(chan bool, 10)
	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		a <- true
		b <- true
		c <- true
	}

	for i := 0; i < 10; i++ {

		select {
		case <-a:
			fmt.Println("a")
		case <-b:
			fmt.Println("b")
		case <-c:
			fmt.Println("c")
		}
		fmt.Println("__")
	}
*/
}