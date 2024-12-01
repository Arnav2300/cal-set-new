package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)
var jwtSecret = []byte("abcdefgh")
func SignJwtToken(role, email, username string) (string, error) {
	 // TODO: Replace with a secure secret from environment variables
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(10 * 24 * time.Hour).Unix(), // Expiry set for 10 days
		"authorized": true,
		"username":   username,
		"email":      email,
		"role":       role,
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err // Return empty string and error if signing fails
	}
	return tokenString, nil // Return token as string and error as nil after signing
}

func VerifyJwtToken(tokenString string) error {
	token,err:=jwt.Parse(tokenString,func(token *jwt.Token) (interface{},error){
		return jwtSecret,nil
	})
	if err!=nil{
		return err
	}
	if !token.Valid{
		return error.New("Invalid token")
	}
	return nil
}