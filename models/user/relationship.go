package user

import (
	"time"

	"github.com/riumat/cinehive-be/models/enums"
)

type Relationship struct {
	ID          int                 `gorm:"primaryKey;autoIncrement"`
	RequesterID int                 `gorm:"not null"`
	ReceiverID  int                 `gorm:"not null"`
	Status      enums.RequestStatus `gorm:"not null"`
	RequestedAt time.Time           `gorm:"autoCreateTime"`
	RespondedAt *time.Time          `gorm:"default:null"`
	Requester   User                `gorm:"foreignKey:RequesterID"`
	Receiver    User                `gorm:"foreignKey:ReceiverID"`
}
