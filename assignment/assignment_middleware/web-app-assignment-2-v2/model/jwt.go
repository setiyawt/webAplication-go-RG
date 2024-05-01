package model

import "github.com/golang-jwt/jwt"

var JwtKey = []byte("session_token")

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
