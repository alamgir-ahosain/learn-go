package main

import (
	"fmt"
	"math"
)
 type shape interface{
	area() float64
 }


type circle struct{
	radius float64
}
func(c circle) area() float64{
	return math.Pi*c.radius*c.radius
}


type rectangle struct {
	height float64
	width float64
}
func(r *rectangle) area() float64 {
	return r.height*r.width
}


//Function that uses the interface — Polymorphic
func get_area(s shape) float64{
	return s.area()
}

func main(){

	/*__Interfaces
	- when share the same method (circle and rectangle have same area() method )
	-used as abstruction:
	- can be used as Polymorphism.
	 function can operate on different types, as long as they share a common behavior (interface).
	- Go doesn’t use implements. It figures it out automatically — this is called structural typing.

	   Does circle have a method area() float64? 
	   if yes ,then automatically call.does not need implement/extends  keyword like java.



	*/

	c1:=circle{4.5}
	r1:=rectangle{2,3}
	// fmt.Println(c1.area());fmt.Println(r1.area())

	shapes:=[]shape{c1,&r1}
	for _,val:=range shapes{
		//fmt.Println(idx.area())
		fmt.Println(get_area(val))
	}

}