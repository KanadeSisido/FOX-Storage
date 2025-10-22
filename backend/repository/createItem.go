package repository

import (
	"context"
	"fox-storage/model"
	"time"

	"gorm.io/gorm"
)

func (r itemRepository) CreateItem(ctx context.Context, name string, typ model.ItemType, parentId *string, ownerId string, mimetype *string, sizeBytes *int64) (model.Item, error) {

	item := model.Item{
		Name:      name,
		Type:      typ,
		ParentID:  parentId,
		OwnerID:   ownerId,
		MimeType:  mimetype,
		SizeBytes: sizeBytes,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {

		err := gorm.G[model.Item](r.db).Create(ctx, &item)

		if err != nil {
			return err
		}

		perm := model.Permission{
			ItemID:     item.ID,
			UserID:     ownerId,
			Permission: "Owner",
		}

		err = gorm.G[model.Permission](r.db).Create(ctx, &perm)

		if err != nil {
			return err
		}

		return nil

	})

	return item, err
}
