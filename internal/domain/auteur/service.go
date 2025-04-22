package auteur

import "github.com/gin-gonic/gin"

type AuteurService interface {
	CreateAuteur(ctx *gin.Context) (Auteur, error)
	GetAuteur() ([]Auteur, error)
	FindById() (Auteur , error)
	UpdateAuteur(id uint , ctx *gin.Context) (Auteur, error)
	DeleteAuteur(id uint) (error)
}