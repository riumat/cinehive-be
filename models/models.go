package models

import (
	"time"
)

type ContentType string

const (
	ContentTypeMovie ContentType = "movie"
	ContentTypeTV    ContentType = "tv"
)

type RequestStatus string

const (
	RequestStatusPending  RequestStatus = "pending"
	RequestStatusAccepted RequestStatus = "accepted"
	RequestStatusRejected RequestStatus = "rejected"
)

// Models

type User struct {
	UserID                 int       `gorm:"primaryKey;autoIncrement"`
	Username               string    `gorm:"size:50;not null"`
	Email                  string    `gorm:"size:100;unique;not null"`
	Password               string    `gorm:"size:255;not null"`
	CreatedAt              time.Time `gorm:"autoCreateTime"`
	Watchtime              int       `gorm:"default:0"`
	FeaturedContentID      *int
	FeaturedContentType    *ContentType
	Contents               []Content      `gorm:"foreignKey:UserID"`
	Persons                []Person       `gorm:"foreignKey:UserID"`
	RequesterRelationships []Relationship `gorm:"foreignKey:RequesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReceiverRelationships  []Relationship `gorm:"foreignKey:ReceiverID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Watchlists             []Watchlist    `gorm:"foreignKey:UserID"`
}

type Content struct {
	UserID      int
	ContentID   int
	ContentType ContentType `gorm:"type:content_type_enum"`
	Rating      *int
	Review      *string   `gorm:"size:400"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Watchlist struct {
	UserID      int
	ContentID   int
	ContentType ContentType `gorm:"type:content_type_enum"`
	ContentName string
	CreatedAt   time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Person struct {
	PersonID int
	UserID   int

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Relationship struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	RequesterID int
	ReceiverID  int
	Status      RequestStatus `gorm:"type:request_status_enum"`
	RequestedAt time.Time     `gorm:"autoCreateTime"`
	RespondedAt *time.Time

	Requester User `gorm:"foreignKey:RequesterID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Receiver  User `gorm:"foreignKey:ReceiverID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
