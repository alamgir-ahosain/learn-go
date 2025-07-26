package main

import (
	"fmt"
)

type myError struct{
	errorMsg string
}

//implementing the error interface
func (my myError) Error() string { //Error()->override the bultin Error()
	return my.errorMsg
}

func custom_msg() error{
      return myError{"cannot devide by zero."}
}


func divideByZero(x ,y int)(int ,error){
	if y==0 {
		return -1,custom_msg() 
	}else{
		return x/y,nil
	}

}
func main(){

	//handle divide by zero
	res,err:=divideByZero(10,0)
	if err!=nil{
		fmt.Println("Error:",err.Error())
		fmt.Println("res:",res)
	}else{
		fmt.Println(res)
	}

	
}