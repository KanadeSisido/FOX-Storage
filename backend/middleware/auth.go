package middleware

import (
	"errors"
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

	fmt.Println(tokenStr)
	
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


	switch {
		case token.Valid:
			fmt.Println("You look nice today")		
		case errors.Is(err, jwt.ErrTokenMalformed):
			fmt.Println("That's not even a token")
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			fmt.Println("Invalid signature")
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		default:
			fmt.Println("Couldn't handle this token:", err)
	}

}