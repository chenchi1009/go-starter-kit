package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Rules    []Rule `json:"rules" gorm:"many2many:user_rules;"`
}

func (u *User) TableName() string {
	return "sys_users"
}
