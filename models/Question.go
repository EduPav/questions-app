package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model

	Description string `json:"description"`
	Rating      uint   `gorm:"default:1" json:"rating"`
	CreatorID   uint   `json:"creator_id"`
}
