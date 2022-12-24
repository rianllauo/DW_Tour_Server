package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountry() ([]models.Country, error)
	GetCountry(ID int) (models.Country, error)
	CreateCountry(country models.Country) (models.Country, error)
	// UpdateUser(user models.User) (models.User, error)
	DeleteCountry(country models.Country) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCountry() ([]models.Country, error) {
	var Countries []models.Country
	err := r.db.Find(&Countries).Error

	return Countries, err
}

func (r *repository) GetCountry(ID int) (models.Country, error) {
	var country models.Country
	err := r.db.First(&country, ID).Error

	return country, err
}

func (r *repository) CreateCountry(Country models.Country) (models.Country, error) {
	err := r.db.Create(&Country).Error

	return Country, err
}

// func (r *repository) UpdateUser(user models.User) (models.User, error) {
// 	err := r.db.Save(&user).Error

// 	return user, err
// }

func (r *repository) DeleteCountry(country models.Country) (models.Country, error) {
	err := r.db.Delete(&country).Error

	return country, err
}
