package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemType string

const (
    ItemFile   ItemType = "file"
    ItemFolder ItemType = "folder"
)

type Item struct {
    ID        string     `gorm:"type:char(36);primaryKey"`
    Name      string     `gorm:"size:255;not null"`
    Type      ItemType   `gorm:"type:enum('file','folder');not null"`
    ParentID  *string    `gorm:"type:char(36);index"`
    OwnerID   string     `gorm:"type:char(36);not null;index"`
    MimeType  *string    `gorm:"size:100"`
    SizeBytes *int64
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
    if i.ID == "" {
        i.ID = uuid.New().String()
    }
    return
}
