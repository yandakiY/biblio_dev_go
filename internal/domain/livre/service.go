package livre

type LivreService interface {
	GetLivre() []Livre
	FindById(id uint) (Livre, error)
	CreateLivre(livre *Livre) (Livre, error)
	UpdateLivre(id uint, livre *Livre) (Livre, error)
	DeleteLivre(id uint)
}