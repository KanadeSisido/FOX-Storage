package repository

import (
	"context"
	"fox-storage/model"

	"gorm.io/gorm"
)

func (r itemRepository) GetItems(ctx context.Context, parentId string) ([]model.Item, error) {

	items, err := gorm.G[model.Item](r.db).Where("parent_id = ?", parentId).Order("FIELD(type, 'folder', 'file')").Find(ctx)

	if err != nil {
		return nil, err
	}

	return items, nil
}
