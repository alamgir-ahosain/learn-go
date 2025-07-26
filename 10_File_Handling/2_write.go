package main

import (
	"bufio"
	"fmt"
	"os"
)

var fn = fmt.Println

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//! create a file
	f, err1 := os.Create("create_demo.txt")
	check_error(err1)
	defer f.Close() //close the file

	//write in  file
	f.WriteString("write in file\n")
	f.WriteString("CSE,MBSTU\n")
	bytes := []byte("using slice\n")
	f.Write(bytes)

    //file read
	data, err2 := os.ReadFile("create_demo.txt")
	check_error(err2)
	fn(string(data))


	//!___________ read and write to another file(streaming fashion)
	//read from create_demo file and write in demo 

	sourceFile,err3:=os.Open("create_demo.txt")
	check_error(err3)
	defer sourceFile.Close()
/*
	
	destFile,err3:=os.Open("demo.txt") //read only
	os.O_WRONLY : open 
	os.CREATE :create if does not exist
	os.TRUNC: truncate if already exist
*/
destFile, err3 := os.OpenFile("demo.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	check_error(err3)
	defer destFile.Close()

	reader:=bufio.NewReader(sourceFile)
	writer:=bufio.NewWriter(destFile)

	for{
		 b,err:=reader.ReadByte()
		 if err!=nil{
			if err.Error()!="EOF"{
				panic(err)
			}
			 break
		 }
		 err2:=writer.WriteByte(b)
		 if err2!=nil{
			panic(err2)
		 }
	}
	writer.Flush()
	fn("Written File Succesfully!")

	//!delete file 
	err4:=os.Remove("file.txt")
	if err4!=nil{
		panic(err4)
	}else{
		fn("file delete succesfully!")
	}
	















}
