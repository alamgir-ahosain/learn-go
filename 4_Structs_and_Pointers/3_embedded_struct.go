package main

import (
	"fmt"
)

// Embedded Struct
type user struct{
	Name string
	Id int
}

type info struct{
	Dept string
	Choice bool
	user // embedded (anonymous) field

}

func main(){

	//inline way
	inf:=info{
		Dept:"CSE",
		Choice: true,
		user:user{
			Name: "alamgir",
			Id: 12,
		},
	}
	fmt.Println(inf)
	inf.user.Name="Alamgir";fmt.Println(inf)

	//or this way
	user_me:=user{Name: "alamgir",Id: 12,}
	inf2:=info{Dept:"CSE",Choice: true,user:user_me,}
	fmt.Println(inf2)
 
	


}