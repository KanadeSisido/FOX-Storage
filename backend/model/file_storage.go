package model

import (
	"time"
)

type FileStorage struct {
	ItemID     string  `gorm:"type:char(36);primaryKey"`
	StorageKey string  `gorm:"size:255;not null"`
	Checksum   *string `gorm:"size:128"`
	UploadedAt time.Time
}

type FileItemLocation struct {
	ItemID     string  `gorm:"type:char(36);primaryKey"`
	StorageKey string  `gorm:"size:255;not null"`
	Checksum   *string `gorm:"size:128"`
	UploadedAt time.Time
	ID         string   `gorm:"type:char(36);primaryKey"`
	Name       string   `gorm:"size:255;not null"`
	Type       ItemType `gorm:"type:enum('file','folder');not null"`
	ParentID   *string  `gorm:"type:char(36);index"`
	OwnerID    string   `gorm:"type:char(36);not null;index"`
	MimeType   *string  `gorm:"size:100"`
	SizeBytes  *int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
