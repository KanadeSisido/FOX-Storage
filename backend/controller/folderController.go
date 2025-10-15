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
		var mime string;

		if item.SizeBytes == nil {
			sizeBytes = 0
		} else {
			sizeBytes = *item.SizeBytes
		}

		if item.MimeType == nil {
			mime = ""
		} else {
			mime = *item.MimeType
		}

		itemDtos = append(itemDtos, dto.ItemsDto{
			ID: item.ID,
			Name: item.Name,
			Type: item.Type,
			SizeBytes: sizeBytes,
			Mime: mime,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return itemDtos
}