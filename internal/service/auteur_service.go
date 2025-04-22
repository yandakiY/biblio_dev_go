package service

import (
	"github.com/yandakiY/biblio_dev_go/internal/domain/auteur"
	"github.com/yandakiY/biblio_dev_go/internal/repository"
)

type auteurService struct {
	repo repository.AuteurRepository
}

type AuteurService interface {
	GetAuteur() []auteur.Auteur
	FindById(id uint) *auteur.Auteur
	CreateAuteur(a *auteur.Auteur) (*auteur.Auteur, error)
	UpdateAuteur(id uint, a *auteur.Auteur) (*auteur.Auteur, error)
	DeleteAuteur(id uint) error
}

func NewAuteurService(repo repository.AuteurRepository) AuteurService {
	return &auteurService{
		repo: repo,
	}
}


// CreateAuteur implements AuteurService.
func (service *auteurService) CreateAuteur(a *auteur.Auteur) (*auteur.Auteur, error) {
	return service.repo.CreateAuteur(a)
}

// DeleteAuteur implements AuteurService.
func (a *auteurService) DeleteAuteur(id uint) error {
	return a.repo.DeleteAuteur(id)
}

// FindById implements AuteurService.
func (a *auteurService) FindById(id uint) *auteur.Auteur {
	return a.repo.FindById(id)
}

// GetAuteur implements AuteurService.
func (a *auteurService) GetAuteur() []auteur.Auteur {
	return a.repo.GetAuteur()
}

// UpdateAuteur implements AuteurService.
func (service *auteurService) UpdateAuteur(id uint, a *auteur.Auteur) (*auteur.Auteur, error) {
	return service.repo.UpdateAuteur(id ,a)
}
