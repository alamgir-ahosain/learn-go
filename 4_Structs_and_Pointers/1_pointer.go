package main

import (
	"fmt"
)

func show(a[3] int){
	a[0]=12
}
func show_ref(a *[3] int){
	//*a[0]=12 cannot indirect a[0] (variable of type int)
	(*a)[0]=12  //deference the array then index it
}

//a *[3] int ->pointer to an array of 3 integer
//a[3] *int ->an array of 3 pinter 


func main(){
	
	/*____Pointer
	store the memory address
	* ->deference operator,get the value 
	*/

	arr:=[3] int{1,2,3} ;fmt.Println(arr)
	show(arr)  //pass by value
	fmt.Println("After pass by value: ",arr)
	show_ref(&arr) //pass by reference
	fmt.Println("After pass by reference",arr)

     x:=12
	 p:=&x // p is now pointing the memory address of x
	 *p=53 // modify p=53 in the address of p -> ultimatly the address of x
	 fmt.Println(x,p,*p)  //53 0xc00000a110 53

}