package main

import (
	"fmt"
)
func main(){
	
	/*
	i. capacity:maximul number of elements can fit
	ii. underlying array:if modify one sclice then effect the others that share the same array
	iii. dynamic array
	*/

	//empty or nil slice
	var s[] int 
	fmt.Println("slice:",s," Length:",len(s)," capacity:",cap(s))

	//Make function with length
	 s2:=make([]int,3) //type,size
	 fmt.Println("slice:",s2," Length:",len(s2)," capacity:",cap(s2))
	 	
	 //Make function with length and capacity
	 s3:=make([]int,3,10) //type,size,capacity . appent up to 7 more element
	 fmt.Println("slice:",s3," Length:",len(s3)," capacity:",cap(s3))

	 //append operation
	 s3=append(s3,1,2,3)
	 fmt.Println("slice:",s3," Length:",len(s3)," capacity:",cap(s3))

	 //slice literal:way to initialize
	 s4:=[]int{0,1,2,3,4,5,6,7,8,9} ;fmt.Println("slice:",s4," Length:",len(s4)," capacity:",cap(s4))

	 //slice from a slice: use slice operator
	 a:=s4[2:5];fmt.Println("slice a:",a)
	 b:=s4[2:] ;fmt.Println("slice b:",b)
	 c:=s4[:6] ;fmt.Println("slice c:",c)
	 d:=make([]int,len(c)); copy(d,c); fmt.Println("slice d:",d) 

    //slice from an array
	arr:=[]int{0,1,2,3,4};fmt.Println("array:",arr)
	f:=arr[1:4] ;fmt.Println("slice f:",f)

    //slice of slice
	matrix := [][]int{ {1, 2}, {3, 4},} ;fmt.Println(matrix)
	matrix = append(matrix, []int{1, 2})
	matrix = append(matrix, []int{3, 4})
	fmt.Println(matrix)

	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, val)
		}
		fmt.Println()
	}

	//underlynig array
	array:=[5] int{0,1,2,3,4} 
	x:=array[1:4]    //1 2 3
	y:=x[1:]         //2,3
	fmt.Println(array);	fmt.Println(x);	fmt.Println(y);fmt.Println()

	//effect other 
	// slice does not copy value ,if referencing other 
	x[1]=12
	fmt.Println(array);	fmt.Println(x);	fmt.Println(y);fmt.Println()


}