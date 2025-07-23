package main

import (
	"fmt"
)
func main(){

	// classic for loop
	for i:=1;i<5;i++{
		fmt.Print(i," ")
	}
	fmt.Println()


	//infinite loop
			// for i:=1;;{
			// 	fmt.Println(i)
			// }


	//* Range	
	//  1 to 5 
	for i:=range 5{
		fmt.Print(i," ")
	}		
	fmt.Println()


	//continue and break
	for i:=range 12{
		if i==5{
			continue //skip the iteration
		}
		if i==11{
             break  //exit the loop
		}

		fmt.Print(i," ")
	}
	fmt.Println()


	//!GO has only one loop: it's  mighty for loop
	// so implemet while loop:using for loop
	i:=1
	for i<5{
		fmt.Print(i," ")
		i++
	}


	

}