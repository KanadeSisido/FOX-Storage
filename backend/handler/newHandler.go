package handler

import (
	"fox-storage/controller"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	LoginHandler(ctx *gin.Context)
	RootNameHandler(ctx *gin.Context)
	SignUpHandler(ctx *gin.Context)
}

type userHandler struct {
	controller controller.UserController
}

func NewUserHandler(_controller controller.UserController) UserHandler {

	handler := userHandler{controller: _controller}

	return handler
}

type ItemHandler interface {
	FolderHandler(ctx *gin.Context)
	FileHandler(ctx *gin.Context)
	CreateFolderHandler(ctx *gin.Context)
	CreateFileHandler(ctx *gin.Context)
}

type itemHandler struct {
	controller controller.ItemController
}

func NewItemHandler(_controller controller.ItemController) ItemHandler {

	handler := itemHandler{controller: _controller}

	return handler
}
