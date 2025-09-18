package dto

import (
	"fox-storage/model"
	"time"
)


type ItemsDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Type      model.ItemType `json:"type"`
	SizeBytes int64     `json:"size"`
	UpdatedAt time.Time `json:"updated_at"`
}