package model

import (
	"time"
)

type FileStorage struct {
    ItemID     string    `gorm:"type:char(36);primaryKey"`
    StorageKey string    `gorm:"size:255;not null"`
    Checksum   *string   `gorm:"size:128"`
    UploadedAt time.Time
}
