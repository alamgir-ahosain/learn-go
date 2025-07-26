package user

import (
	"fmt"
)

     type UserInfo  struct {
		Name string
		Dept string
	 }

     //capital letter,can access outside the package
	 func PrintUser(u UserInfo) {
		  fmt.Println("Name:",u.Name," Dept:",u.Dept)
	 }

	 


