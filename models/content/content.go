package content

import (
	"time"

	"github.com/riumat/cinehive-be/models/enums"
	"github.com/riumat/cinehive-be/models/user"
)

type Content struct {
	UserID         int               `gorm:"primaryKey"`
	ContentID      int               `gorm:"primaryKey"`
	ContentType    enums.ContentType `gorm:"primaryKey"`
	Rating         *int              `gorm:"default:null"`
	Review         *string           `gorm:"size:400"`
	CreatedAt      time.Time         `gorm:"autoCreateTime"`
	User           user.User         `gorm:"foreignKey:UserID"`
	ContentToGenre []ContentToGenre  `gorm:"foreignKey:UserID,ContentID,ContentType"`
}
