package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    ID        string         `gorm:"type:char(36);primaryKey"`
    Email     string         `gorm:"size:255;uniqueIndex;not null"`
    Name      string         `gorm:"size:255;not null"`
    RootID    string         `gorm:"type:char(36)"`
	HashedPassword string
    CreatedAt time.Time
    UpdatedAt time.Time
}

// 新規作成時にUUIDを自動生成
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    if u.ID == "" {
        u.ID = uuid.New().String()
    }
    return
}
