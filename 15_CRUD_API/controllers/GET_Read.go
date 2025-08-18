package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/alamgir-ahosain/learn-go/tree/main/15_CRUD_API/model"
)

func GET_ReadOperation() {

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	//check getting responce
	if response.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status:", response.Status)
		return
	}
	//  Read data: Generic Responce
	// 	var todo Todo
	// 	err = json.NewDecoder(responce.Body).Decode(&todo)
	// 	if err != nil {
	// 		fmt.Println("Error decoading:", err)
	// 		return
	// 	}
	// 	fmt.Println(todo)

	// }
	var todo model.Todo
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Todo:", todo)
}

