package service

import (
	"fmt"
	"fox-storage/middleware"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s userService) CreateJwtToken(userId string) (*string, error) {

	claims := middleware.Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "FOX-Storage",
		},
	}

	keyData, err := os.ReadFile("./config/private_key.pem")

	if err != nil {
		fmt.Println("Error: Private key cannot Open")
		return nil, err
	}

	privateKeyStr, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)

	if err != nil {
		fmt.Println("Error: Private key cannot Parse")
		return nil, err
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	_token, err := unsignedToken.SignedString(privateKeyStr)

	if err != nil {
		fmt.Println("Error: cannot create token")
		return nil, err
	}

	return &_token, nil
}
