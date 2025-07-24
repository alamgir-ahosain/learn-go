package main

import (
	"fmt"
)

func add(x,y int) int {
	return x+y
	
}

func pass_func(x,y int, p func(a,b int) int) int{
      return p(x,y)
}

func return_func(name string) func(int) string {
	return func(id int)string {
		return fmt.Sprintf("Name: %s ,ID:%d",name,id)
	}
}


func main(){

	/*___higher order function if___
	pass function as parameter.
	return a function.
	or both
	*/

	//!pass a function
	// callback function:passing a function  in another function as argument
	result:=pass_func(1,2,add)
	fmt.Println(result)

	//!return a function
	result2:=return_func("alamgir") 
	//result =func(id int)string{return fmt.Sprintf("Name: %s ,ID:%d",name,id) }
	fmt.Println(result2(12))











}