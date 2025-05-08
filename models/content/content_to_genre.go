package content

import (
	"time"

	"github.com/riumat/cinehive-be/models/enums"
	"github.com/riumat/cinehive-be/models/user"
)

type ContentToGenre struct {
	ID           int               `gorm:"primaryKey;autoIncrement"`
	UserID       int               `gorm:"not null"`
	ContentID    int               `gorm:"not null"`
	ContentType  enums.ContentType `gorm:"not null"`
	GenreID      int               `gorm:"not null"`
	CreatedAt    time.Time         `gorm:"autoCreateTime"`
	User         user.User         `gorm:"foreignKey:UserID"`
	ContentGenre ContentGenre      `gorm:"foreignKey:GenreID"`
	Content      Content           `gorm:"foreignKey:UserID,ContentID,ContentType"`
}
