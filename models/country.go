package models

type Country struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" gorm:"type: varchar(255)"`
	// Trip TripResponse `json:"trip"`
	// CreatedAt time.Time `json:"-"`
	// UpdatedAt time.Time `json:"-"`
}

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Trip TripResponse `json:"trip"`
}

func (CountryResponse) TableName() string {
	return "countries"
}
