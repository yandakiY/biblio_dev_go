package livre

import "github.com/yandakiY/biblio_dev_go/internal/domain/auteur"

type LivreGetList struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type LivreGetListDetails struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Auteur      auteur.Auteur `json:"auteur"`
}

type LivreCreateNew struct {
	Name string `json:"name" binding:"required" validate:"required"`
	Description string `json:"description" binding:"required" validate:"required"`
	AuteurId uint `json:"auteur_id" validate:"required" binding:"required"`
}

type LivreUpdate struct {
	Name string `json:"name"`
	Description string `json:"description"`
	AuteurId uint `json:"auteur_id"`
}