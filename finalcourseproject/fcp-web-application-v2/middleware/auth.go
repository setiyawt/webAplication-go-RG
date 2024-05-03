package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil {
			ctx.Redirect(http.StatusSeeOther, "/login")
			ctx.Abort()
			return
		}

		tokenString := cookie.Value
		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			// Menyimpan UserID dari claims ke dalam context
			ctx.Set("email", claims.Email)
			// Melanjutkan ke middleware/handler selanjutnya
			ctx.Next()
		} else {

			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		}
	})
}
