package content

type ContentGenre struct {
	ID             int              `gorm:"primaryKey"`
	Name           string           `gorm:"not null"`
	ContentToGenre []ContentToGenre `gorm:"foreignKey:GenreID"`
}
