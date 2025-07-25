package main

import (
	"fmt"
)
func print_int(slc []int){
        for _,val:=range slc{
			fmt.Print(val," ")
		} 
		fmt.Println()
}
func print_string(slc []string){
        for _,val:=range slc{
			fmt.Print(val," ")
		} 
		fmt.Println()
}

	/*
	[T any ]( slc [] T) ->print any type 
	[T interface{}]( slc [] T)->print any type
	[T int | string ]( slc [] T) ->only int and string
	[T comparable]( slc [] T) ->type of interface,comparable type
	*/
func use_generics[T any]( slc [] T){
	 for _,val:=range slc{
			fmt.Print(val," ")
		} 
		fmt.Println()
}
func main(){

      /*
	 print_int() and print_string() are same function signature.that's create duplication.
	 use generics to avoid code duplication.
	  */
	x:=[] int{1,2,3}
	y:=[] string{"alamgir","CSE","MBSTU"}
	print_int(x); print_string(y) //
	 
	use_generics(x); use_generics(y)


	

}