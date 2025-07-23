package main

import (
	"fmt"
)
func main(){

	var x int32=12; fmt.Printf("int number:%d\n",x) //positive and neagitive
	var y uint32=13;fmt.Println(y) //only positive

	var x2 float32=12.543; fmt.Printf("Float number:%0.2f\n",x2)
	var str string="alamgir" ;fmt.Printf("String :%s\n",str)

	//! ________rune :alias for int32 bit._________________
	//     store only unicode format
	// %c character, %U unicode format
	var r rune='😊' ;
	fmt.Printf("Emoji as character format:%c\n",r)
	fmt.Printf("Emoji as unicode:%U\n",r)

	name:="alamgir"
	id:=12
	happy:=true


    //%v : any type,auto detect
	fmt.Printf("name:%v\n",name) //string
	fmt.Printf("id:%v\n",id)   //int
	fmt.Printf("happy:%v\n",happy) //bool
	
	//* complex Data type
	 c1:=1+2i; c2:=complex(3,2)
	 fmt.Println("c1=",c1,",c2=",c2,",c1+c2",c1+c2)
	 fmt.Printf("c1 = %v, c2 = %v, c1 + c2 = %v\n", c1, c2, c1 + c2)




		








}
