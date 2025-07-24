package main

import (
	"fmt"
)
func sum(val ...int){
	 s:=0
	 //range val return two value-index,value
	 for _,i:=range val{
		s+=i
	 }
	 fmt.Print(val," ");
	 fmt.Print(" sum:",s);fmt.Println()


}
func main(){
	/*___variadic function
	 take zero or more argument of same type.
	 ...(ellipsis  operator).
	 function argument treated as slice.
	*/
	sum(1,2,9)
	slice:=[]int {1,2,10}
	sum(slice...) //slice 

}