package controller

import "context"

func (c *Controller) GetFileController(ctx context.Context, fileId string, userId string) (*string, *string, error) {

	return c.service.GetFileName(ctx, fileId, userId)
}