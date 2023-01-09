package models

import "time"

type Image struct{}

type Trip struct {
	ID             int             `json:"id" gorm:"primary_key:auto_increment"`
	Title          string          `json:"name" form:"title" gorm:"type: varchar(255)" validate:"required"`
	CountryID      int             `json:"country_id"  `
	Country        CountryResponse `json:"country" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Accomodation   string          `json:"accomodation" form:"accomodation" gorm:"type: varchar(255)" validate:"required"`
	Transportation string          `json:"transportation" form:"transportation" gorm:"type: varchar(255)"`
	Eat            string          `json:"eat" form:"eat" gorm:"type: varchar(255)"`
	Day            int             `json:"day" form:"day"`
	Night          int             `json:"night" form:"night"`
	DateTrip       time.Time       `json:"date_trip" form:"datetrip"`
	Price          int             `json:"price" form:"price"`
	Quota          int             `json:"quota" form:"quota"`
	Description    string          `json:"description" form:"description"`
	Image          []string        `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserId         int             `json:"user_id"`
}

type TripResponse struct {
	ID             int             `json:"id"`
	Title          string          `json:"name" `
	CountryID      int             `json:"country_id"  `
	Country        CountryResponse `json:"country" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Accomodation   string          `json:"accomodation"`
	Transportation string          `json:"transportation"`
	Eat            string          `json:"eat" `
	Day            int             `json:"day" `
	Night          int             `json:"night"`
	DateTrip       time.Time       `json:"date_trip" `
	Price          int             `json:"price" `
	Quota          int             `json:"quota" `
	Description    string          `json:"description" `
	Image          []string        `json:"image"`
	UserId         int             `json:"user_id"`
}

func (TripResponse) TableName() string {
	return "trips"
}
