package main

import (
	"fmt"
	"sync"
)

// Receivig ONLY
func receving_Chan(ch <-chan int, wg *sync.WaitGroup) {
	val, is_Channel_Open := <-ch
	if is_Channel_Open {
		fmt.Println(val)
	} else {
		fmt.Println("channel is closed")
		fmt.Println("default val:", val)
	}
	//fmt.Println(<-myChan)
	wg.Done()
}

// send ONLY
func send_chan(ch chan<- int, wg *sync.WaitGroup) {

	//close(myChan) -> panic :send on closed channel
	ch <- 5
	ch <- 10
	close(ch) //ok,pass the value and then close the channel.

	wg.Done()
}

func main() {

    /*___channels
	synchronized communication pipes .
	safely pass data between goroutines.
	*/

	/*
		Error: all goroutines are asleep - deadlock!
		myChan:=make(chan int)
		myChan <- 5
		fmt.Println(<-myChan)
	*/

	myChan := make(chan int, 2) //string buffering upto 2 values
	var wg sync.WaitGroup

	wg.Add(2) //two goroutine
	go receving_Chan(myChan, &wg)
	go send_chan(myChan, &wg)

	wg.Wait()
}
