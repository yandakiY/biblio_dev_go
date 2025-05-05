package auteur

import "github.com/yandakiY/biblio_dev_go/internal/domain/dto/livre"

type AuteurGetList struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	// Livres   []livre.LivreGetList `json:"livres"`
}

type AuteurGetListDetails struct{
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Livres   []livre.LivreGetList `json:"livres"`
}

type AuteurCreateNew struct{
	Name     string `json:"name" validate:"required" binding:"required"`
	LastName string `json:"last_name" validate:"required" binding:"required"`
}

type AuteurUpdate struct{
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}