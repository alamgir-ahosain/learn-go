package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Users struct {
		Name string `json:"name"`
		SID  string `jsonL:"sid"`
	}
	user := Users{
		Name: "Alamgir",
		SID:  "CE21012",
	}
	fmt.Println(user) //{Alamgir CE21012}

	//convert user into json, Encoading(Marshalling)
	data, err := json.Marshal(user) // byte format
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data)) //{"name":"Alamgir","SID":"CE21012"}

	//json into user,json Decoading(Unmarshalling)
	var userData Users
	err = json.Unmarshal(data, &userData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userData) //{Alamgir CE21012}

}
