package trip

import "time"

type TripRequest struct {
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	// Country        string    `json:"country" gorm:"many2many:trip_country"`
	CountryID      int       `json:"country_id" form:"country_id" gorm:"type: int"`
	Accomodation   string    `json:"accomodation" form:"accomodation" gorm:"type: varchar(255)"`
	Transportation string    `json:"transportation" form:"transportation" gorm:"type: varchar(255)"`
	Eat            string    `json:"eat" form:"eat" gorm:"type: varchar(255)"`
	Day            int       `json:"day" form:"day"`
	Night          int       `json:"night" form:"night"`
	DateTrip       time.Time `json:"date_trip" form:"date_trip"`
	Price          int       `json:"price" form:"price"`
	Quota          int       `json:"quota" form:"quota"`
	Description    string    `json:"description" form:"description"`
	Image          []string    `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type UpdateTripRequest struct {
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	// Country        string    `json:"country" gorm:"many2many:trip_country"`
	CountryID      int       `json:"country_id" form:"country_id" gorm:"type: int"`
	Accomodation   string    `json:"accomodation" form:"accomodation" gorm:"type: varchar(255)"`
	Transportation string    `json:"transportation" form:"transportation" gorm:"type: varchar(255)"`
	Eat            string    `json:"eat" form:"eat" gorm:"type: varchar(255)"`
	Day            int       `json:"day" form:"day"`
	Night          int       `json:"night" form:"night"`
	DateTrip       time.Time `json:"date_trip" form:"date_trip"`
	Price          int       `json:"price" form:"price"`
	Quota          int       `json:"quota" form:"quota"`
	Description    string    `json:"description" form:"description"`
	Image          string    `json:"image" form:"image" gorm:"type: varchar(255)"`
}
