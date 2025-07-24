package main

import (
	"fmt"
)
func rertun_multiple() (string,int){
return "alamgir",12
}

func add_two(a,b int){ //parameter a,b
	fmt.Println("sum:",a+b)
}


func main(){


	add_two(10,2) //argument 10,2
	
	//!Ruturn Multiple Value
	name,id:=rertun_multiple()
	fmt.Println("name:",name,"\nID:",id)

	//name2,id2:=rertun_multiple()    //if use only name2 then an error occured:id2 is declared but not used
	name2,_:=rertun_multiple()       //use  underscore sign(_) ,blank identifier,ignore the value
	fmt.Println("name:",name2) 


}

