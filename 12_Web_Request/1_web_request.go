package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Web Request..")
	responce, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") //Free fake and reliable API for testing and prototyping.
	defer responce.Body.Close()                                               //close the responce body when main exit

	data, err := ioutil.ReadAll(responce.Body) //data as byte format
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(data)) //convert byte to string
}
