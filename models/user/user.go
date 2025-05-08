package user

import (
	"time"

	"github.com/riumat/cinehive-be/models/common"
	"github.com/riumat/cinehive-be/models/enums"
)

type User struct {
	UserID              int                     `gorm:"primaryKey;autoIncrement;column:user_id"`
	Username            string                  `gorm:"size:50;not null"`
	Email               string                  `gorm:"size:100;unique;not null"`
	Password            string                  `gorm:"size:255;not null"`
	CreatedAt           time.Time               `gorm:"autoCreateTime;column:created_at"`
	Watchtime           int                     `gorm:"default:0"`
	FeaturedContentID   *int                    `gorm:"column:featured_content_id"`
	FeaturedContentType *enums.ContentType      `gorm:"column:featured_content_type"`
	Content             []common.Content        `gorm:"foreignKey:UserID"`
	Person              []Person                `gorm:"foreignKey:UserID"`
	Requester           []Relationship          `gorm:"foreignKey:RequesterID"`
	Receiver            []Relationship          `gorm:"foreignKey:ReceiverID"`
	Watchlist           []Watchlist             `gorm:"foreignKey:UserID"`
	ContentToGenre      []common.ContentToGenre `gorm:"foreignKey:UserID"`
}
