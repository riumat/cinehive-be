package user

type Person struct {
	PersonID int  `gorm:"primaryKey"`
	UserID   int  `gorm:"primaryKey"`
	User     User `gorm:"foreignKey:UserID"`
}
