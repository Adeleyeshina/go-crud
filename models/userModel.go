package models

import (
	"gorm.io/gorm"
)

type User struct {
	//ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      string `json:"age"`
}
