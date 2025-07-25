package main

import (
	"fmt"
)

// enumrated type 
type status int
const (
	_ status = iota //start with 0
	Running
	Failed
	Pending
	Success
)

// enumrated type 
type status2 string
const(
	Received  status2="Received"
	Confirmed status2 ="Confirmed"
	Delivered status2="Delivered"
)
func change_status(s status2){
	fmt.Println("status2:",s)
}


func main() {
	/*
	-GO does not have enum type
	 but can be implement with const and iota
    -special type that represnet a group of constant value.

	*/

	var dd status = Running;  fmt.Println(dd)
	change_status(Confirmed)

}
