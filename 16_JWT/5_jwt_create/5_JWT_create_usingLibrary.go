package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*	steps for JWT creation
	1. Header -> Base64 (b64Header)
	2. Payload -> Base64 (b64Payload)
	3. Combine: headerPayload := b64Header + "." + b64Payload
	4. Signature: HMAC-SHA256(secretKey, headerPayload)
	5. Encode signature -> Base64 (hash)
	6. JWT:  jwt := b64Header + "." + b64Payload + "." + hash

	-> output token check : jwt.io

*/

// JWT Playload
type Payload struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool
}

// Create JWT
func CreateJWT(secretKey string, payload Payload) (string, error) {
	claims := jwt.MapClaims{
		"id":            payload.ID,
		"first_name":    payload.FirstName,
		"last_name":     payload.LastName,
		"email":         payload.Email,
		"is_shop_owner": payload.IsShopOwner,
		"exp":           time.Now().Add(time.Hour * 24).Unix(), // token expires in 24h
	}

	// Create token with HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	return token.SignedString([]byte(secretKey))
}

func main() {
	secret := "my-secret"
	playload := Payload{
		ID:          12,
		FirstName:   "Alamgir",
		LastName:    "Hosain",
		Email:       "alamgir@gmail.com",
		IsShopOwner: false,
	}

	token, err := CreateJWT(secret, playload)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)

}
