package controller

import (
	"context"
	"fox-storage/dto"
	"fox-storage/model"
)

func (c *Controller) FolderController(ctx context.Context, folderId string, userId string) (*[]dto.ItemsDto, error) {
	items, err := c.service.GetItems(ctx, folderId, userId)

	if err != nil {
		return nil, err
	}

	itemDtos := itemsToItemsDto(items)
	
	return &itemDtos, nil
}


func itemsToItemsDto(items []model.Item) []dto.ItemsDto {
	
	var itemDtos []dto.ItemsDto;


	for _, item := range items {

		var sizeBytes int64

		if item.SizeBytes == nil {
			sizeBytes = 0
		} else {
			sizeBytes = *item.SizeBytes
		}

		itemDtos = append(itemDtos, dto.ItemsDto{
			ID: item.ID,
			Name: item.Name,
			Type: item.Type,
			SizeBytes: sizeBytes,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return itemDtos
}