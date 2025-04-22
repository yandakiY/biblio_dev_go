package service

import (
	"github.com/yandakiY/biblio_dev_go/internal/domain/livre"
	"github.com/yandakiY/biblio_dev_go/internal/repository"
)

type livreService struct {
	repo repository.LivreRepository
}


type LivreService interface {
	GetLivre() []livre.Livre
	FindById(id uint) (*livre.Livre, error)
	CreateLivre(livre *livre.Livre) (*livre.Livre, error)
	UpdateLivre(livre *livre.Livre) (*livre.Livre, error)
	DeleteLivre(id uint) error
}

func NewLivreService(repo repository.LivreRepository) LivreService {
	return &livreService{
		repo: repo,
	}
}


// CreateLivre implements LivreService.
func (l *livreService) CreateLivre(livre *livre.Livre) (*livre.Livre, error) {
	return l.repo.CreateLivre(livre)
}

// DeleteLivre implements LivreService.
func (l *livreService) DeleteLivre(id uint) error {
	return l.repo.DeleteLivre(id)
}

// FindById implements LivreService.
func (l *livreService) FindById(id uint) (*livre.Livre, error) {
	return l.repo.FindById(id)
}

// GetLivre implements LivreService.
func (l *livreService) GetLivre() []livre.Livre {
	return l.repo.GetLivre()
}

// UpdateLivre implements LivreService.
func (l *livreService) UpdateLivre(livre *livre.Livre) (*livre.Livre, error) {
	return l.repo.UpdateLivre(livre)
}
