package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	ParentID uint   `json:"parent_id" gorm:"default:0"`
	Icon     string `json:"icon"`
	Path     string `json:"path" gorm:"unique;not null"`
	Order    int    `json:"order"`
}

func (m *Menu) TableName() string {
	return "sys_menus"
}
