package main

import (
	"fmt"
)
func main(){


	//!
	 /*
	  ______format_______
		if {
		}else if{
		}else{
		}

		________wrong way____
		if {
		}
		else if{
		}
		else {
		}

	 */

	x:=12
	if x<=5 {
		fmt.Println("less than or equal 5")
	}else if x!=10{                            
		fmt.Println("number is not equal 10")
	}else{
		fmt.Println("YOU know knothing John Snow!")
	}


	//!declared variable inside 
	if y:=5;y==5{
		fmt.Println("value is 5")
	}else{
		fmt.Println("error")
	}

	//!	go does not have ternary,only normal if else
	//int x = (a > b) ? a : b;  // Ternary




}