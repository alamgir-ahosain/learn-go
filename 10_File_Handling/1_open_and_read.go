package main

import (
	"fmt"
	"os"
)

var fn=fmt.Println


 func check_error(err error){
	if err!=nil{
		//fmt.Println("Error->",err)
		panic(err)
	}
 }
func main(){

	//!______ Open File
	f,err1:=os.Open("demo.txt") //open and read only
	defer f.Close() //close when main function is complete
	check_error(err1)

	fileInfo,err2:=f.Stat() //fil info
	check_error(err2)

	//get File Info
	fn("File Name:",fileInfo.Name())
	fn("File or Folder:",fileInfo.IsDir())
	fn("File Size",fileInfo.Size())
	fn("File Permission:",fileInfo.Mode()) //on my disk
	fn("File Modified at:",fileInfo.ModTime())


	//!_______Read File 
	
	// *1.one and simplest way :slow

	data,err4:=os.ReadFile("deo.txt") 
	check_error(err4)
	fn(string(data))

	// *2. another way
	buffer:=make([]byte,fileInfo.Size())  //buffer: temporary space,array of byte.
	_,err3:=f.Read(buffer)  //d,err3:=f.Read(buffer) //d=buffer size
	check_error(err3)
	for i:=0;i<len(buffer);i++{
            fn("Data",string(buffer[i]))
	}

	//!__ ___Read Directory (folder and file)
	dir,err5:=os.Open(".") //current folder, (../)previous folder
	check_error(err5)

	//12: show at most 12 file and folder, -1:show all files and folder
	info,err6:=dir.ReadDir(-1) 
	check_error(err6)
	for _,fi:=range info{
		fn(fi.Name())
	}
	

}