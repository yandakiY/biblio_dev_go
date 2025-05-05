package auteur

import (
	"github.com/yandakiY/biblio_dev_go/internal/domain/livre"
	"gorm.io/gorm"
)

type Auteur struct {
	gorm.Model
	Name string `json:"name" binding:"required" validate:"required"`
	LastName string `json:"last_name" binding:"required" validate:"required"`
	Livres []livre.Livre `json:"livres";gorm:foreignKey:"AuteurId"` 
}