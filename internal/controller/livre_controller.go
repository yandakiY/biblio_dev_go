package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yandakiY/biblio_dev_go/internal/domain/livre"
	"github.com/yandakiY/biblio_dev_go/internal/service"
)

type livreController struct {
	service service.LivreService
}


type LivreController interface {
	Create(ctx *gin.Context) (*livre.Livre, error)
	Update(ctx *gin.Context) (*livre.Livre, error)
	Get() []livre.Livre
	FindById(id uint) *livre.Livre
	Delete(id uint) error
}

func NewLivreController(service service.LivreService) LivreController {
	return &livreController{
		service: service,
	}
}

// Create implements LivreController.
func (l *livreController) Create(ctx *gin.Context) (*livre.Livre, error) {
	var livre livre.Livre

	if err := ctx.ShouldBindJSON(&livre); err != nil{
		return nil , err
	}

	res , err1 := l.service.CreateLivre(&livre)
	if err1 != nil{
		return nil , err1
	}

	return res, nil
}

// Delete implements LivreController.
func (l *livreController) Delete(id uint) error {
	if err := l.service.DeleteLivre(id); err != nil{
		return err
	}
	return nil
}

// FindById implements LivreController.
func (l *livreController) FindById(id uint) *livre.Livre {
	
	res , err := l.service.FindById(id)
	if err != nil {
		return nil
	}
	return res
}

// Get implements LivreController.
func (l *livreController) Get() []livre.Livre {
	return l.service.GetLivre()
}

// Update implements LivreController.
func (l *livreController) Update(ctx *gin.Context) (*livre.Livre, error) {
	
	var livre livre.Livre
	if err := ctx.ShouldBindJSON(livre); err != nil {
		return nil, err
	}

	res , err := l.service.UpdateLivre(&livre)
	if err != nil {
		return nil , err
	}

	return res , nil
}
