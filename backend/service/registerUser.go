package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s userService) RegisterUser(ctx context.Context, username string, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return err
	}

	return s.repository.CreateUser(ctx, username, email, string(hashedPassword))
}
