package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h itemHandler) FolderHandler(ctx *gin.Context) {

	parentID := ctx.Param("folderId")
	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c := ctx.Request.Context()
	data, err := h.controller.FolderController(c, parentID, userId.(string))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"items": data,
	})
}
