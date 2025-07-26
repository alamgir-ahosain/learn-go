package main

import (
	
	"github.com/alamgir-ahosain/learn-go/operation"
	"github.com/alamgir-ahosain/learn-go/user"
)
func main(){

	operation.Additon(12,3)
	
	inf:=user.UserInfo{
		Name: "alamgir",
		Dept: "CSE",
	}
       
	user.PrintUser(inf)

}