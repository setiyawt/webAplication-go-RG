package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Mengecek apakah terdapat cookie session_token dalam request
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil || cookie.Value == "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0LCJleHAiOjE3MTQ0ODM0MTB9.9zau65zvdlMdlvNYa8fVI-SKWnwFFNU-JCXLEB0zNMo" {
			// Jika tidak ada cookie session_token, kirim respon HTTP 401 Unauthorized
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}

		// Mengekstrak token JWT dari cookie session_token
		tokenString := cookie.Value
		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil // Menggunakan kunci JWT yang sama yang digunakan untuk signing
		})
		if err != nil {
			// Jika token tidak valid, kirim respon HTTP 400 Bad Request
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// Mengecek apakah token memiliki claims yang sesuai
		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			// Menyimpan UserID dari claims ke dalam context
			ctx.Set("id", claims.UserID)
			// Melanjutkan ke middleware/handler selanjutnya
			ctx.Next()
		} else {
			// Jika token tidak valid atau tidak memiliki claims yang sesuai, kirim respon HTTP 400 Bad Request
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		}
	})
}
