package main

import (
	"fmt"
)
func main(){

	//basic switch case
	x:=12
	switch x{
	case 1,2,3:
		fmt.Println("1 to 3")
	default:
		fmt.Println("default case")
	}

	//! fallthrough: Always falls to the next case
    y:=12
	switch y{
	case 12:
		fmt.Println("is 12")
		fallthrough
	case 13:
		fmt.Println("is 13")
	default:
		fmt.Println("default case")
	}

	/*
	output:
			is 12
			is 13
	*/


	//switch case :like if -else if -else
	    z:=12
	switch {
	case z!=12:
		fmt.Println("not  12")
	case z==12:
		fmt.Println("is 12")
	default:
		fmt.Println("default case")
	}
	

	//*type switch: check type of variable in runtime
	//interface{}:accept any value
	variable_type:=func(i interface{}){
		switch i:=i.(type) {
		case int:
			fmt.Println(i,":type is int")
		case string:
			fmt.Println(i,":type is string")
		default:
			fmt.Println("default case")
			
		}
	}
	variable_type(12)
	variable_type("alamgir")





}