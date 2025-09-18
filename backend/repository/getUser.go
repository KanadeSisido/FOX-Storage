package repository

import (
	"context"
	"fox-storage/model"

	"gorm.io/gorm"
)

func (r *Repository) GetUserByUsername(ctx *context.Context, username string) (*model.User, error) {
	user, err := gorm.G[model.User](r.db).Where("name = ?", username).First(*ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserById(ctx *context.Context, userId string) (*model.User, error) {
	user, err := gorm.G[model.User](r.db).Where("id = ?", userId).First(*ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

