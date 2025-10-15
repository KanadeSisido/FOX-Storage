package controller

import (
	"context"
	"io"
)

func (c *Controller) UploadController(ctx *context.Context, filename string, file io.Reader, bytesSize int64, mime string, parentId string, userId string) error {
	return c.service.CreateFile(ctx, filename, file, bytesSize, mime, &parentId, userId)
}