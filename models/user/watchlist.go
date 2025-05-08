package user

import (
	"time"

	"github.com/riumat/cinehive-be/models/enums"
)

type Watchlist struct {
	UserID      int               `gorm:"primaryKey"`
	ContentID   int               `gorm:"primaryKey"`
	ContentType enums.ContentType `gorm:"primaryKey"`
	ContentName string            `gorm:"not null"`
	CreatedAt   time.Time         `gorm:"autoCreateTime"`
	User        User              `gorm:"foreignKey:UserID"`
}
