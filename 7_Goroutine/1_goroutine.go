package main

import (
	"fmt"
	"time"
)

var cnt=0
func task(i int){
	cnt=i
	fmt.Print(i," ")
}
func main(){
	/*
	    - run functions concurrently.
		- Managed by go scheduler not OS.
		- go routine stack size is 2 kb.
	*/

	/*____Race Condition___
	in the end,Final cnt Not guaranteed to be 5
	two or more goroutines (or threads) access shared data at the same time
	-> (Race condition->solve with channel,sync.Mutex (Lock/Unlock)

	*/
	for i:=0;i<6;i++{
		// go task(i)
		func(ii int){
			go task(ii)
		}(i)
	}
	
	time.Sleep(time.Second*2)
	fmt.Println("\n Final cnt:",cnt)
}
