package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


type Claims struct {
	UserID string `json:"UserID"`;
	jwt.RegisteredClaims
}


func Auth(c *gin.Context) {

	tokenStr, err := c.Cookie("Authorization")
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Un Authorized Error"})
		return
	}

	keyData, err := os.ReadFile("config/private_key.pem")

	if err != nil {
		fmt.Println("Error: Private key cannot Open")
		return
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)

	if err != nil {
		fmt.Println("Error: cannot parse PEM")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return privateKey.Public(), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
		return
	}

	c.Set("userId", claims.UserID)

}