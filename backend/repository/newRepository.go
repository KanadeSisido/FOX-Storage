package repository

import (
	"context"
	"fox-storage/model"
	"io"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(ctx context.Context, name string, typ model.ItemType, parentId *string, ownerId string, mimetype *string, sizeBytes *int64) (model.Item, error)
	GetItems(ctx context.Context, parentId string) ([]model.Item, error)
	GetItemLocationById(ctx context.Context, fileId string) (*model.FileItemLocation, error)
	GetPermissionByUserID(ctx context.Context, itemId string, userId string) (*model.Permission, error)
	LocateFile(ctx context.Context, id string, fileName string, file io.Reader) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return itemRepository{db: db}
}

type UserRepository interface {
	CreateUser(ctx context.Context, username string, email string, hashedPassword string) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}
