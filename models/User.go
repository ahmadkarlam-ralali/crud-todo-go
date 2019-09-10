package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username	string	`gorm:"column:username" json:"username"`
	Password	string	`gorm:"column:password" json:"password"`
}
