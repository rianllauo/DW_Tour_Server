package countryDto

type CountryRequest struct {
	Name string `json:"name" form:"name" gorm:"type: varchar(255)"`
}
