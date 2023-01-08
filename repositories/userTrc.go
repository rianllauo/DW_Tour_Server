package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type UserTrcRepository interface {
	FindUserTrc(ID int) ([]models.Transaction, error)
	// GetUser(ID int) (models.User, error)
	// CreateUser(user models.User) (models.User, error)
	// UpdateUser(user models.User) (models.User, error)
	// DeleteUser(user models.User) (models.User, error)
}

func RepositoryUserTrc(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUserTrc(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").Where("user_id =?", ID).Find(&transactions).Error

	// err := r.db.Preload("Trip").Raw("SELECT * FROM transactions WHERE user_id=?", ID).Scan(&transactions).Error

	return transactions, err
}
