package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

/*
			SHA-256 = Secure Hash Algorithm 256-bit
	1. Cryptographic hash function
	2. Fixed output: 256 bits (32 bytes/64 hex chars)
	3. One-way (irreversible)
	4. Same input -> always same hash
	5. Small input change -> big output change
	6. Very hard to find two different inputs with same hash
	7. Used in passwords, signatures, integrity, blockchain


*/

func main() {
	data := []byte("alamgir")

	// Create SHA-256 hash
	hash := sha256.Sum256(data)

	//convert hash(byte array) to hex string
	hashString := hex.EncodeToString(hash[:])
	fmt.Println(hash)
	fmt.Println(hashString)
}
