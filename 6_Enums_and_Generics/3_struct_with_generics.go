package main
import (
	"fmt"
)

type user[T any,U any,V any] struct{
	Info[] T
	Name T
	Dept U
	Option V


}
func main(){

	myStack:=user[int,string,bool]{
		Info: [] int{1,2,3},
		Name:12,
		Dept:"CSE",
		Option: true,
	}
	fmt.Println(myStack)


	myStack2:=user[string,int,float64]{
		Info: [] string{"alamgir","hosain"},
		Name:"alamgir",
		Dept:12,
		Option: 3.5,
	}
	fmt.Println(myStack2)
}