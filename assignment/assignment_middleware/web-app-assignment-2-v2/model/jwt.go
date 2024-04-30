package model

import "github.com/golang-jwt/jwt"

var JwtKey = []byte("gtgergfrfewde433fd09jddkfsdsw")

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
