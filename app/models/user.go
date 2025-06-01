package models

type User struct {
	Model
	Username string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null;default:null"`
	Status   int8   `gorm:"default:1"`
	IsAdmin  uint8  `gorm:"default:0"`
}
