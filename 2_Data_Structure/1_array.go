package main

import (
	"fmt"
)
func main(){

	/* _____Basic type__
	default value :int 0 ,string "", bool false
	*/
	var a[3] int; fmt.Println(a)        //declartion without initalization
	b:=[3] int{1,2,3} ;fmt.Println(b)   //declartion with initalization
	c:=[...] int{1,2,3} ;fmt.Println(c) //inferred length
	d:=[3] int{0:10,1:11,2:12} ;fmt.Println(d) //indexed with initalization
	d2:=[5] int{1:11,3:12} ;fmt.Println(d2)     
    e:=[2][3] int{{1,2,3},{4,5,6},};fmt.Println(e) //2D array


	//Print array
	for i:=0;i<len(c);i++{
		fmt.Print(c[i]," ")
	} 

	fmt.Println()
	for idx:=range c{
		fmt.Print(c[idx]," ")
	}
	
	fmt.Print("\n2D array: ")
	for i:=range 2{
		for j:=range 3{
			print(e[i][j]," ")
		}
	}

}