package repository

import (
	"context"
	"fox-storage/model"
)

func (r itemRepository) GetItemLocationById(ctx context.Context, fileId string) (*model.FileItemLocation, error) {
	item := model.FileItemLocation{}

	err := r.db.Table("items").
		Select(`items.id AS id,
                items.name AS name,
                items.type AS type,
                items.parent_id AS parent_id,
                items.owner_id AS owner_id,
                items.mime_type AS mime_type,
                items.size_bytes AS size_bytes,
                items.created_at AS created_at,
                items.updated_at AS updated_at,
                file_storages.item_id AS item_id,
                file_storages.storage_key AS storage_key,
                file_storages.checksum AS checksum,
                file_storages.uploaded_at AS uploaded_at`).
		Joins("JOIN file_storages ON file_storages.item_id = items.id").
		Where("items.id = ?", fileId).
		First(&item).Error

	if err != nil {
		return nil, err
	}
	return &item, nil
}
