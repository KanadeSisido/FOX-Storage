package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PermissionLevel string

const (
    PermissionRead  PermissionLevel = "read"
    PermissionWrite PermissionLevel = "write"
    PermissionOwner PermissionLevel = "owner"
)

type Permission struct {
    ID         string          `gorm:"type:char(36);primaryKey"`
    ItemID     string          `gorm:"type:char(36);not null;index"`
    UserID     string          `gorm:"type:char(36);not null;index"`
    Permission PermissionLevel `gorm:"type:enum('read','write','owner');not null"`
    CreatedAt  time.Time
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
    if p.ID == "" {
        p.ID = uuid.New().String()
    }
    return
}
