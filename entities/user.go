package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	HP       string `json:"hp" form:"hp"`
	Token    string
	Books    []Book `gorm:"foreignKey:Author_ID"`
}
