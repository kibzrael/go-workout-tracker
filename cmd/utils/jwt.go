package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey string = os.Getenv("JWT_PRIVATE_KEY")

func EncodeJWT(payload jwt.MapClaims) (result string, ok error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	result, err := token.SignedString([]byte(JwtKey))
	if err != nil{
		return "", err
	}
	return result, nil
}

func DecodeJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return []byte(JwtKey), nil
	})
	if err != nil{
		return nil, err
	} else if !token.Valid{
		return nil, errors.New("invalid token")
	}
	
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}