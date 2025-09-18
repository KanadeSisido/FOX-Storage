package controller

import (
	"context"
	"io"
)

func (c *Controller) UploadController(ctx *context.Context, filename string, file io.Reader, parentId string, userId string) error {
	return c.service.CreateFile(ctx, filename, file, &parentId, userId);
}