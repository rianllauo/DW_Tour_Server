package models

// User model struct
type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(100)"`
	Address  string `json:"address" gorm:"type: text"`
	Role     string `json:"role" gorm:"type: varchar(100)"`
	Avatar   string `json:"avatar" gorm:"type: text"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(100)"`
	Address  string `json:"address" gorm:"type: text"`
	Avatar   string `json:"avatar" gorm:"type: text"`
}

func (UserResponse) TableName() string {
	return "users"
}
