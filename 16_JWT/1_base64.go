package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := "alamgir"

	// Encode string to Base64
	byteArr := []byte(data)
	encoding := base64.URLEncoding
	encoding = encoding.WithPadding(base64.NoPadding)
	byte64Str := encoding.EncodeToString(byteArr)

	//short way
	encoded := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println("Data :", data)
	fmt.Println("Byte Format: ", byteArr)
	fmt.Println("Base64 Encoded: ", byte64Str)
	fmt.Println("Base64 Encoded: ", encoded)

	// Decode string to Base64
	decoded, err := encoding.DecodeString(byte64Str)
	decoded2, err2 := base64.StdEncoding.DecodeString(encoded)
	if err != nil && err2 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Base64 Decoded: ", decoded)
	fmt.Println("Decoded:", string(decoded))
	fmt.Println("Decoded:", string(decoded2))

}
