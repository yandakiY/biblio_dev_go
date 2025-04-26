package repository

import (
	"fmt"

	"github.com/yandakiY/biblio_dev_go/internal/domain/livre"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type livreRepo struct{
	conn *gorm.DB
}

type LivreRepository interface{
	GetLivre() []livre.Livre
	FindById(id uint) (*livre.Livre, error)
	CreateLivre(livre *livre.Livre) (*livre.Livre , error)
	UpdateLivre(id uint , livre *livre.Livre) (*livre.Livre, error)
	DeleteLivre(id uint) error
}

func NewLivreRepository() LivreRepository{

	conn , err := gorm.Open(sqlite.Open("biblio_dev.db"), &gorm.Config{})

	if err != nil {
		panic("Error while the loading of database...")
	}

	err = conn.AutoMigrate(&livre.Livre{})
	if err != nil{
		panic("Error while the migration of items...")
	}

	return &livreRepo{
		conn: conn,
	}
}

func (repo *livreRepo) GetLivre() []livre.Livre{
	var livres []livre.Livre
	if err := repo.conn.Set("auto_preload", true).Find(&livres).Error; err != nil {
		return nil
	}
	return livres
}

func (repo *livreRepo) FindById(id uint) (*livre.Livre, error){
	var livre *livre.Livre
	if err := repo.conn.Find(&livre, id).Error; err != nil{
		return nil , err
	}
	return livre, nil
}

func (repo *livreRepo) CreateLivre(livre *livre.Livre) (*livre.Livre , error){
	if err := repo.conn.Create(livre).Error; err != nil{
		return nil , err
	}
	return livre, nil
} 

func (repo *livreRepo) UpdateLivre(id uint , l *livre.Livre) (*livre.Livre, error){

	// find the livre with id
	var existing *livre.Livre
	if err := repo.conn.First(&existing , id).Error ; err != nil {
		return nil , err
	}

	if err := repo.conn.Model(&existing).Updates(l).Error ; err != nil {
		return nil , err
	}
	return existing, nil
}

func (repo *livreRepo) DeleteLivre(id uint) error {
	res := repo.conn.Unscoped().Delete(&livre.Livre{}, id)
	
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("aucun livre avec l'id %d", id)
	}
	fmt.Println("Error:", res.Error)
	fmt.Println("RowsAffected:", res.RowsAffected)


	return nil
}