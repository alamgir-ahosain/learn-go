package main

import "fmt"
func main(){

	//Explicit declartion
	var x string;  x="alamgir 1"; fmt.Println(x) 

	//implicit declartion
	var y string="alamgir 2" ;fmt.Println(y) 
	var y1,y2="first","second" ;fmt.Println(y1,y2)

	//group
	var(
		name string="alamgir"
		id int=12

	)
	fmt.Println(name,id)

	//go automatically infer the type
	//only used inside the function,local variable
	z:="alamgir 3" ;fmt.Println(z)
	z1,z2:="first","second"; fmt.Println(z1,z2)
	fmt.Printf("The type of z is: %T\n", z) // T for any type ;means when type is unknown


	//constant_(not data type,used to declare variable)
	const pi=3.14 //untyped string
	const pi2 float64=3.14
	fmt.Println(pi,pi2)


}