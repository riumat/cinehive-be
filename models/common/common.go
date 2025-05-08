package common

type User struct {
	UserID int `gorm:"primaryKey"`
}

type Content struct {
	ContentID   int    `gorm:"primaryKey"`
	ContentType string `gorm:"primaryKey"`
}

type ContentToGenre struct {
	ContentID   int    `gorm:"primaryKey"`
	ContentType string `gorm:"primaryKey"`
	GenreID     int    `gorm:"primaryKey"`
}
