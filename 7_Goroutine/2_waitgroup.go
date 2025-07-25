package main

import (
	"fmt"
	"sync"
	
)

var wg sync.WaitGroup //pointer 

var cnt=0
func task(i int,wg *sync.WaitGroup){
	defer wg.Done()
	cnt=i
	fmt.Print(i," ")
}
func main(){
	
     /*
	 allow main function to wait untill al goroutine are done.
	 wg.Add(1) ->set the number of goroutine.
	 defer wg.Done()  ->tell when this goroutine is done.
	 wg.Wait()->prevents exit main until all goroutine are done.


	 */
	for i:=0;i<6;i++{
		// go task(i)
		func(ii int){
			wg.Add(1) 
			go task(ii,&wg)
			
		}(i)
	}
	wg.Wait()
	fmt.Println("\n Final cnt:",cnt)
	
}
