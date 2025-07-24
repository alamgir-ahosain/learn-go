package main

import (
	"fmt"
)

func counter() func() int{
     cnt:=0
	 return func()int{
		cnt++
		return  cnt
	 }
}
func add(x int) func(y int) int{
	return func(y int) int{
		return x+y
      
	}

}


func main(){
	/*
	remember variable form outer function,even the outer function is closed.
	it can access and use those outer function captured variable.

	use escape analysis:go compiler determine that  value is to  stored or not.
	go compiler knew that vvariable in the function is later used.


	*/

	cnt1:=counter()
	cnt2:=counter()
	fmt.Println(cnt1(),cnt1()) //1 2
	fmt.Println(cnt2(),cnt2()) //1 2


	//proved that:outer function variable will be stored.
	a:=add(5)  //retrun a function and remember x=5
	b:=add(10) //retrun a function and remember x=10
	fmt.Println(a(5)) //10
	fmt.Println(b(2)) //12


}