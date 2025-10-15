package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateFileHandler(ctx *gin.Context) {

	parentId := ctx.Param("folderId")
	file, err := ctx.FormFile("file")
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "no attached file",
		})
		return
	}
	
	userId, exists := ctx.Get("userId")
	
	if !exists {
		ctx.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Unauthorized",
			},
		)
		return
	}

	bytesSize := file.Size
	mime := file.Header.Get("Content-Type")

	f, err := file.Open()
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot open file",
		})
		return
	}

	defer f.Close()

	c := ctx.Request.Context()

	err = h.controller.UploadController(&c, file.Filename, f, bytesSize, mime,parentId, userId.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":"created",
	})
}