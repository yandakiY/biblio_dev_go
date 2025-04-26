package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yandakiY/biblio_dev_go/internal/domain/auteur"
	"github.com/yandakiY/biblio_dev_go/internal/service"
)

type auteurController struct {
	service service.AuteurService
}


type AuteurController interface {
	Create(ctx *gin.Context) (*auteur.Auteur, error)
	Update(id uint, ctx *gin.Context) (*auteur.Auteur, error)
	Get() []auteur.Auteur
	FindById(id uint) (*auteur.Auteur , error)
	Delete(id uint) error
}

func NewAuteurController(service service.AuteurService) AuteurController {
	return &auteurController{
		service: service,
	}
}

// Create implements AuteurController.
func (a *auteurController) Create(ctx *gin.Context) (*auteur.Auteur, error) {
	// bind object to ctx
	var auteur auteur.Auteur
	if err := ctx.ShouldBindJSON(&auteur); err != nil {
		return nil , err
	}

	res , err1 := a.service.CreateAuteur(&auteur)
	if err1 != nil {
		return nil , err1
	}

	return res, nil
}

// Delete implements AuteurController.
func (a *auteurController) Delete(id uint) error {
	if err := a.service.DeleteAuteur(id); err != nil{
		return err
	}
	return nil
}

// FindById implements AuteurController.
func (a *auteurController) FindById(id uint) (*auteur.Auteur , error) {
	auteur , err := a.service.FindById(id)
	if err != nil{
		return nil , err
	}
	return auteur , nil
}

// Get implements AuteurController.
func (a *auteurController) Get() []auteur.Auteur {
	return a.service.GetAuteur()
}

// Update implements AuteurController.
func (a *auteurController) Update(id uint, ctx *gin.Context) (*auteur.Auteur, error) {
	var auteur auteur.Auteur

	if err := ctx.ShouldBindJSON(&auteur); err != nil {
		return nil , err
	}

	res , err1 := a.service.UpdateAuteur(id , &auteur)
	if err1 != nil {
		return nil , err1
	}
	return res, nil
}
