package controller

import "context"

func (c *Controller) RootNameController(ctx *context.Context, userId string) (*string, error) {
	return c.service.GetRootIdByUserId(ctx,userId)
}