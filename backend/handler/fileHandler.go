package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h itemsHandler) FileHandler(ctx *gin.Context) {

	fileId := ctx.Param("id")
	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c := ctx.Request.Context()
	localName, fileName, err := h.controller.GetFileController(c, fileId, userId.(string))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "no such file",
		})
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+*fileName)
	ctx.File(*localName)
}
