package controller

import "context"

func (c itemController) GetFileController(ctx context.Context, fileId string, userId string) (*string, *string, error) {

	return c.service.GetFileName(ctx, fileId, userId)
}
