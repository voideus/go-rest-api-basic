package middlewares

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/voideus/go-rest-api/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"message": "No auth header passed.",
			})
			c.Abort()
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		if tokenString == "" {
			c.JSON(401, gin.H{
				"message": "No token string passed.",
			})
			return
		}

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
