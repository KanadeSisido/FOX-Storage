package controller

import "context"

func (c itemController) CreateFolderController(ctx context.Context, name string, parentId *string, userId string) error {
	return c.service.CreateFolder(ctx, name, parentId, userId)
}
