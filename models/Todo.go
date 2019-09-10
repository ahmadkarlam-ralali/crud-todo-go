package models

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	UserId uint `gorm:"column:user_id" json:"user_id"`
	Title  string `gorm:"column:title" json:"title"`
	Status int `gorm:"column:status;size:10" json:"status"`
}
