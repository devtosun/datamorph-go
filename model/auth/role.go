package model

import "gorm.io/gorm"

type Role struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
	gorm.Model
}
