package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

/*
		HMAC (Hash-based Message Authentication Code)
		
	1. Uses a secret key + hash function (e.g., SHA-256).
	2. Output length = depends on hash function (e.g., 256 bits for SHA-256)
	3. Provides integrity (data not altered)
	4. Provides authenticity (from trusted sender)
	5. Resistant to tampering and length-extension attacks
	6. Used in APIs, JWTs, secure tokens, SSL/TLS, etc.



*/

func main() {
	message := "alamgir"
	secretKey := "mySecretKey"

	//Create HMAC with SHA-256
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(message))

	// Final HMAC value
	mac := hash.Sum(nil)

	fmt.Println("MAC value:", mac)
	fmt.Println("HMAC-SHA256: ", hex.EncodeToString(mac))

}
