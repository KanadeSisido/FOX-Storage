package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) RootNameHandler(ctx *gin.Context){

	userId, exists := ctx.Get("userId")

	fmt.Println(userId)

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	c := ctx.Request.Context()
	rootId, err := h.controller.RootNameController(&c, userId.(string))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"rootId": rootId,
	})

}