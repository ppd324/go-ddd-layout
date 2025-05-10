package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"unique;column:nickname;not null"`
	Email    string `gorm:"column:email;not null"`
	Avatar   string `gorm:"column:avatar;not null"`
	Status   int    `gorm:"column:status;not null"` // 用户状态
}
