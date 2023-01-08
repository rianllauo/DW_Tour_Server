package models

type Transaction struct {
	ID         int          `json:"id" gorm:"primary_key:auto_increment"`
	CounterQty int          `json:"counter_qty"`
	Total      int          `json:"total"`
	Status     string       `json:"status"`
	Attachment string       `json:"attachment"`
	TripID     int          `json:"trip_id"`
	Trip       TripResponse `json:"trip" `
	UserId     int          `json:"userId"`
	User       UserResponse `json:"user"`
}

// type TransactionResponse struct {
// 	ID         int          `json:"id" gorm:"primary_key:auto_increment"`
// 	CounterQty int          `json:"counter_qty"`
// 	Total      int          `json:"total"`
// 	Status     string       `json:"status"`
// 	Attachment string       `json:"attachment"`
// 	TripID     int          `json:"transaction_id"`
// 	Trip       TripResponse `json:"transaction" `
// }

// func (TransactionResponse) TableName() string {
// 	return "transactions"
// }
