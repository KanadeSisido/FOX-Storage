package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateFolderDto struct {
	Name string `json:"name"`
}

func (h *Handler) CreateFolderHandler(ctx *gin.Context){

	ParentID := ctx.Param("folderId")
	
	var createFolder CreateFolderDto;
	err := ctx.ShouldBind(&createFolder)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Type",
		})
		return
	}

	userID, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Cannot get your id",
		})
		return
	}

	context := ctx.Request.Context()
	
	err = h.controller.CreateFolderController(&context, createFolder.Name, &ParentID, userID.(string))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "created successfully",	
	})

}