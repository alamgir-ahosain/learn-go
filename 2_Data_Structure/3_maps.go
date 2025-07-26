package main

import (
	"fmt"
	"maps"
)
func main(){

	//KEY and Value pair
	//maps -> hash or object or dict

	mp1:=make(map[string] string)
	mp1["name"]="alamgir" ; mp1["Dept"]="CSE" ; mp1["id"]="CE21012"

	fmt.Println(mp1["name"],mp1["id"])  //get 
	delete(mp1,"name")  //delete
	fmt.Println(mp1) //randomly print 


	mp2:=make(map[int] string)
	mp2[1]="Shajib sir"
	mp2[2]="Mustakim and Badhon"
	mp2[3]="Nahid and Badhon"
	mp2[12]="Alamgir"

	for key,val:=range mp2{
		fmt.Printf("key: %v ,value: %v\n",key,val) 
	}

    //check if a key is exist
	val,ok:=mp2[13]
	if ok{
		fmt.Println("Key exist and value:",val)
	}else{
		fmt.Println("key Not found")
	}

	//check equality
	m1:=map[string]int{"alamgir":12,"hosain":13}
	m2:=map[string]int{"alamgir":12,"jk":13}
	fmt.Println(maps.Equal(m1,m2))

}