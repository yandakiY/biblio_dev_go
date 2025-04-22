package auteur

import "gorm.io/gorm"

type Auteur struct {
	gorm.Model
	Name string `json:"name" binding:"required" validate:"required"`
	LastName string `json:"last_name" binding:"required" validate:"required"`
}