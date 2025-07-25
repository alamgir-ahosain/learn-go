package main

import (
	"fmt"
	"sync"
)
type task struct{
	Task_id int
	mu sync.Mutex
}
func(t *task) inc(wg *sync.WaitGroup){
	
	//wg.Done()
	//t.mu.Unloock() or 
	defer func(){
		t.mu.Unlock()
		wg.Done()
		
	}()
	t.mu.Lock()
	t.Task_id++
	
}

func main(){
	/*
		race condition-> in the end,Final Task_id Not guaranteed to be 12  
		Mutex(Mutual Exclusion)->solve the race condition.
		Lock()->only only goroutine enter 
		Unlock->release lock for next goroutine

	*/

	var wg sync.WaitGroup

	my_task:=task{Task_id: 0}
	for i:=0;i<12;i++{
		wg.Add(1)
		go my_task.inc(&wg)
	}
	wg.Wait()
	fmt.Println("Final Task_id:",my_task.Task_id)
}