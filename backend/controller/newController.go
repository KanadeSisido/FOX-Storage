package controller

import (
	"context"
	"fox-storage/dto"
	"fox-storage/service"
	"io"
)

type UserController interface {
	LoginController(ctx context.Context, username string, password string) (*string, error)
	SignupController(ctx context.Context, username string, email string, password string) error
	RootNameController(ctx context.Context, userId string) (*string, error)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return userController{service: service}
}

type ItemController interface {
	GetFileController(ctx context.Context, fileId string, userId string) (*string, *string, error)
	FolderController(ctx context.Context, folderId string, userId string) (*[]dto.ItemsDto, error)
	CreateFolderController(ctx context.Context, name string, parentId *string, userId string) error
	UploadController(ctx context.Context, filename string, file io.Reader, bytesSize int64, mime string, parentId string, userId string) error
}

type itemController struct {
	service service.ItemService
}

func NewItemController(service service.ItemService) ItemController {
	return itemController{service: service}
}
