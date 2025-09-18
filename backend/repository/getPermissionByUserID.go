package repository

import (
	"context"
	"fox-storage/model"

	"gorm.io/gorm"
)

func (r *Repository) GetPermissionByUserID(ctx *context.Context, itemId string, userId string) (*model.Permission, error) {
	permission, err := gorm.G[model.Permission](r.db).Where("item_id = ? AND user_id = ?", itemId, userId).First(*ctx)

	if err != nil {
		return nil, err
	}

	return &permission, nil
}