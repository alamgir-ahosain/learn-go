package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
func main(){
	
	//! _____Scan and Scanln____
	fmt.Print("Enter x:");var x string;fmt.Scan(&x) //Reads only first word
	fmt.Print("Enter y:");var y string;fmt.Scanln(&y) //until new line
	fmt.Println();fmt.Println("x:",x); fmt.Println("y:",y)
    /*
	______Left over input error____
	   Enter x:alamgir hosain
	   reads x=alamgir, and _hosain_ store in the input buffer.
	   Scanln does not wait for new user input.it see hosain in the input buffer with a newline.
	   so y=hosain ;automaticalyy store
	*/


	//!____Scanf_______
	fmt.Print("enter string: ")
	var name,name2 string
	fmt.Scanf("%s",&name) //formatted string,only one 
	fmt.Scanf(" %[^\n]",&name2) //with space untill new line
	fmt.Printf("name:%s",name);fmt.Printf(" name2:%s",name2);fmt.Println()
     

	//!____with space ______

	fmt.Println("using bufio")
	reader := bufio.NewReader(os.Stdin)

	// Flush any leftover newline from previous Scanf/Scan
	reader.ReadString('\n') // discard rest of line

	fmt.Print("Enter text: ")
	a, _ := reader.ReadString('\n'); a = strings.TrimSpace(a)

	fmt.Print("Enter again: ")
	b, _ := reader.ReadString('\n');b = strings.TrimSpace(b)

	fmt.Println("a:", a);fmt.Println("b:", b)

  



}