package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) TryAuthorization(ctx *context.Context, username string, password string) (*string, error) {

	user, err := s.repository.GetUserByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	
	if err != nil {
		return nil, err
	}

	return &user.ID, nil
}