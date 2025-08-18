package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alamgir-ahosain/learn-go/tree/main/15_CRUD_API/model"
)

func POST_CreateOperation() {
	todo := model.Todo{
		UserID:    12,
		Title:     "Alamgir",
		Completed: false,
	}
	u := "https://jsonplaceholder.typicode.com/todos/"

	//send data in json format
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	//convert string data to reader
	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	//send POST request
	responce, err := http.Post(u, "application/json", jsonReader)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := io.ReadAll(responce.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

}
