package middlewares

import (
	"fmt"
	"net/http"

	"webservice/bd"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var mySigningKey = []byte("my_secret_key")

/*ChequeoBD es el middlew que me permite conocer el estado de la BD */
func ChequeoBD() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bd.CheckConnect() == 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Conexión perdida con la Base de Datos",
			})
			return
		}
		c.Next()
	}
}

func validateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user", claims["user"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}
