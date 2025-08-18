package main

import (
	"fmt"
	"net/url"
)

func main() {

	u := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Printf("Type of URL is :%T\n", u) //string

	parsedUrl, err := url.Parse(u) //convert string to url type
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", parsedUrl)             //*url.URL
	fmt.Println("Scheme :", parsedUrl.Scheme) // https
	fmt.Println("Host: ", parsedUrl.Host)     //jsonplaceholder.typicode.com
	fmt.Println("Path :", parsedUrl.Path)     // /todos/1
	fmt.Println("RawQuery :", parsedUrl.RawQuery)

	//Modifying URL components
	parsedUrl.Path = "/alamgir/1"

	//Create new URL string
	newURL := parsedUrl.String()
	fmt.Println("New URL:", newURL)
}
