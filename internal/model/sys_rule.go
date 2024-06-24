package model

import "gorm.io/gorm"

type Rule struct {
	gorm.Model
	Name  string `json:"name" gorm:"unique;not null"`
	Menus []Menu `json:"menus" gorm:"many2many:rule_menus;"`
}

func (r *Rule) TableName() string {
	return "sys_rules"
}
