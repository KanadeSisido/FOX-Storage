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

func NewUserHandler(_controller controller.UserController) userHandler {

	handler := userHandler{controller: _controller}

	return handler
}

type ItemsHandler interface {
	FolderHandler(ctx *gin.Context)
	FileHandler(ctx *gin.Context)
	CreateFolderHandler(ctx *gin.Context)
	CreateFileHandler(ctx *gin.Context)
}

type itemsHandler struct {
	controller controller.ItemController
}

func NewHandler(_controller controller.ItemController) itemsHandler {

	handler := itemsHandler{controller: _controller}

	return handler
}
