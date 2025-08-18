package controllers

import (
	"fmt"
	"io"
	"net/http"
)

func DELETE_DeleteOperation() {

	const u = "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest("DELETE", u, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	// Read response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response Body:", string(body))

}
