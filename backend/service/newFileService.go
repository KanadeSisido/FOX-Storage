package service

import (
	"context"
	"fox-storage/model"
	"fox-storage/repository"
	"io"
)

type ItemService interface {
	CreateFile(ctx context.Context, name string, file io.Reader, bytesSize int64, mime string, parentId *string, userId string) error
	CreateFolder(ctx context.Context, name string, parentId *string, userId string) error
	GetFileName(ctx context.Context, fileId string, userId string) (*string, *string, error)
	GetItems(ctx context.Context, folderId string, userId string) ([]model.Item, error)
}

type itemService struct {
	repository repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) (ItemService, error) {
	return &itemService{repository: repo}, nil
}

type UserService interface {
	RegisterUser(ctx context.Context, username string, email string, password string) error
	TryAuthorization(ctx context.Context, username string, password string) (*string, error)
	GetRootIdByUserId(ctx context.Context, userId string) (*string, error)
	CreateJwtToken(userId string) (*string, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) (UserService, error) {
	return &userService{repository: repo}, nil
}
