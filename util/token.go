package util

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// generateJWT takes payload and signs a jwt token containing that info
func GenerateJWT(userId string, email string) (string, error) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:  email,
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	return "Bearer " + tokenString, err
}

// ValidateToken takes a token and returns error if token is not valid
// func ValidateToken(signedToken string) (err error) {
// 	config, err := LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("cannot load config:", err)
// 	}
// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&JWTClaim{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(config.JWTSecret)
// 		},
// 	)
// 	if err != nil {
// 		return
// 	}
// 	claims, ok := token.Claims.(*JWTClaim)
// 	if !ok {
// 		err = error.New("couldn't parse claims")
// 		return
// 	}
// 	if claims.ExpiresAt < time.Now().Local().Unix() {
// 		err = error.New("token expired")
// 		return
// 	}
// 	return
// }
