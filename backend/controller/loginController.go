package controller

import (
	"context"
	"errors"
)

func (c *Controller) LoginController(ctx *context.Context, username string, password string) (*string, error) {

	userId, err := c.service.TryAuthorization(ctx, username, password)

	if err != nil {
		
		return nil, err
	}

	if userId == nil  {
		return nil, errors.New("InvalidUsernameOrPassword")
	}
	
	// token作成
	token, err := c.service.CreateJwtToken(*userId)

	if err != nil {
		return nil, err
	}

	return token, nil

}