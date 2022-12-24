package models

// User model struct
type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(100)"`
	Address  string `json:"address" gorm:"type: text"`
	// TransactionID int                 `json:"transaction_id"`
	// Transaction   TransactionResponse `json:"transaction"`
}

// type UserResponse struct {
// 	ID       int    `json:"id"`
// 	FullName string `json:"name" gorm:"type: varchar(255)"`
// 	Email    string `json:"email" gorm:"type: varchar(255)"`
// }

// func (UserResponse) TableName() string {
// 	return "users"
// }
