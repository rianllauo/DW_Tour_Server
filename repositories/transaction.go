package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
	CreateTransaction(transactions models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, transaction models.Transaction) error
	ChangeTransaction(transaction models.Transaction) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").Preload("User").Find(&transactions, "id = ?", ID).Error

	// err := r.db.Raw("SELECT * FROM transactions WHERE user_id=?", ID).Scan(&transactions).Error

	return transactions, err
}

func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.User").Preload("User").First(&transactions, ID).Error

	return transactions, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip").Preload("Trip.User").Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, transaction models.Transaction) error {
	if status != transaction.Status && status == "success" {
		var trip models.Trip
		r.db.First(&trip, transaction.Trip.ID)
		trip.Quota = trip.Quota - 1
		r.db.Model(&trip).Updates(trip)
	}
	transaction.Status = status
	err := r.db.Preload("Trip").Preload("Trip.User").Preload("User").Model(&transaction).Updates(transaction).Error

	return err
}

func (r *repository) ChangeTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Model(&transaction).Updates(transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
