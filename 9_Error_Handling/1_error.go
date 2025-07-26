package main

import (
	"errors"
	"fmt"
)

func divideByZero(x ,y int)(int ,error){
	if y==0 {
		return -1,errors.New("cannot divide by Zero")//! errors.New() ->create a new error with static message
	}else{
		return x/y,nil
	}

}
func main(){


	//basic error handling
	var a,b int
	it,scanErr:=fmt.Scan(&a,&b) // it store the total number of scanned items
	if scanErr!=nil{
		fmt.Println("Error:",scanErr)
	}else{
		fmt.Println("Scanned items:",it)
		fmt.Println("a:",a," b:",b)
	}


	//handle divide by zero
	res,err:=divideByZero(10,0)
	if err!=nil{
		fmt.Println("Error:",err)
		fmt.Println("res:",res)
	}else{
		fmt.Println(res)
	}

	
}