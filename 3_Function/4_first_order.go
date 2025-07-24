package main

import (
	"fmt"
)
func add(a,b int)int {
	return a+b
}
func multiplication(x int) func(int) int{
            return func(y int) int{
				return x*y
			}
}
func subtraction(x ,y int, f func(int,int) int) int{
	return f(x,y)

}
func main(){

	/*_____First Order Function____
	   1. Name/standerd function
	   2.Anonymous /function literals
	   3.IIFE
	   4.Function Expression
	*/
     
	//_____Function Expression
	//for local scope ,variable must be assign before function invoked.
	// add(1,2) undefined
	sum:=add(1,2);fmt.Println(sum)	//! ____Name/standerd function


	//____anonymus Function:no name
	result:=func(a,b int) int{
		return a+b
	}
	fmt.Println(result(3,4))

	//___IIFE- Immediate Invoked Function Expression
	func(a,b int){
		sum:=a+b
		fmt.Println(sum)
	}(3,4) //invoked 


	//return anonymus function
	m:=multiplication(12)
	fmt.Println(m(10))

	//pass anonymus function
	sub:=subtraction(10,5 ,func(x ,y int ) int{
          return x-y
	})
	fmt.Println(sub)

	
}