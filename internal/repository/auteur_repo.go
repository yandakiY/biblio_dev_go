package repository

import (
	"fmt"

	"github.com/yandakiY/biblio_dev_go/internal/domain/auteur"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type auteurRepo struct {
	conn *gorm.DB
}

type AuteurRepository interface{
	GetAuteur() []auteur.Auteur
	FindById(id uint) (*auteur.Auteur, error)
	CreateAuteur(a *auteur.Auteur) (*auteur.Auteur, error)
	UpdateAuteur(id uint,a *auteur.Auteur) (*auteur.Auteur, error)
	DeleteAuteur(id uint) (error)
}

func NewAuteurRepostitory() AuteurRepository{

	conn , err := gorm.Open(sqlite.Open("biblio_dev.db"), &gorm.Config{})

	if err != nil {
		panic("Error while the loading of database...")
	}

	err = conn.AutoMigrate(&auteur.Auteur{})
	if err != nil {
		panic("Error while the migration of items...")
	}

	return &auteurRepo{
		conn: conn,
	}
}


func (repo *auteurRepo) GetAuteur() []auteur.Auteur{
	var auteurs []auteur.Auteur
	if err := repo.conn.Set("auto_preload", true).Find(&auteurs).Error ; err != nil{
		return nil
	}
	return auteurs
}

func (repo *auteurRepo) FindById(id uint) (*auteur.Auteur , error){
	var auteur *auteur.Auteur
	if err := repo.conn.Find(&auteur, id).Error ; err != nil{
		return nil , err
	}

	if auteur.ID == 0{
		return nil , fmt.Errorf("Auteur introuvable")
	}

	return auteur , nil
}

func (repo *auteurRepo) CreateAuteur(a *auteur.Auteur) (*auteur.Auteur, error){
	if err := repo.conn.Create(a).Error ; err != nil{
		return nil, err
	}
	return a , nil
}

func (repo *auteurRepo) UpdateAuteur(id uint, a *auteur.Auteur) (*auteur.Auteur, error){

	var existing *auteur.Auteur
	// find the auteur with this id
	if err := repo.conn.Find(&existing , id).Error ; err != nil {
		return nil , err
	}

	if existing.ID == 0 {
		return nil , fmt.Errorf("Auteur a modifier inexistant")
	}

	if err := repo.conn.Model(&existing).Updates(a).Error; err != nil{
		return nil, err
	}
	return existing, nil
}

func (repo *auteurRepo) DeleteAuteur(id uint) (error){
	
	res := repo.conn.Unscoped().Delete(&auteur.Auteur{} ,id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("Auteur inexistant")
	}
	return nil
}