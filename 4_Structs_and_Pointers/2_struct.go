package main

import (
	"fmt"
)
type user struct{ // Name,Dept,Id -> called field or property
	Name string
	Dept string
	Id int
}



//! ____Reciever Function/Method

func (u user) get_name() string{
	return u.Name
}

func(u *user)set_name(){
	u.Name="alamgir"
}

//!___Constructor
func user_constructor(name string,dept string,id int) *user{
	user_me:=user{
		 Name:name,
		 Dept:dept,
		 Id:id,
	}
	return &user_me
}

func main(){

	/*
	only copy
	getter ,setter

	struct automaticalyy deferenceing
	A custor data type,collection of different data type

	*/

	var u1 user  //instance /object of struct
	u1=user{ Name: "alamgir",Dept: "CSE"};fmt.Println(u1)
	//u1=user{ "alamgir","CSE",12};fmt.Println(u1)

	//setter
	u1.Name="jk" ;fmt.Println(u1)
	u1.set_name();fmt.Println(u1)

	//getter
	name:=u1.Name;fmt.Println(name) 
	fmt.Println(u1.get_name())
	
	//constructor
	//u1=user_constructor("hosain","CSE",13);fmt.Println(u1)  //cannot use *user as user in assignment
	u2:=user_constructor("hosain","CSE",13);fmt.Println(u2)

	
	// Anonymous Struct->if use only
	info:=struct{
		name string
		id int
	}{"alamgir",12}
	fmt.Println(info)

}
