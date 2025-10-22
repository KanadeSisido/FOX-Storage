package controller

import (
	"context"
)

func (c userController) SignupController(ctx context.Context, username string, email string, password string) error {
	return c.service.RegisterUser(ctx, username, email, password)
}
