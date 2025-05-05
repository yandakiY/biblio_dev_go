package livre

import "gorm.io/gorm"

type Livre struct {
	gorm.Model
	Name string `json:"name" binding:"required" validate:"required"`
	Description string `json:"description" binding:"required" validate:"required"`
	AuteurId uint `json:"auteur_id"`
}