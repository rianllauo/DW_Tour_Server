package models

type Transaction struct {
	ID         int          `json:"id" gorm:"primary_key:auto_increment"`
	CounterQty int          `json:"counter_qty"`
	Total      int          `json:"total"`
	Status     string       `json:"status"`
	Attachment string       `json:"attachment"`
	TripID     int          `json:"product_id"`
	Trip       TripResponse `json:"product"`
	// BuyerID    int          `json:"buyer_id"`
	// Buyer      UserResponse `json:"buyer"`
	// SellerID   int                  `json:"seller_id"`
	// Seller     UsersProfileResponse `json:"seller"`
	// Price      int                  `json:"price"`
	// Status    string               `json:"status"  gorm:"type:varchar(25)"`
	// CreatedAt time.Time `json:"-"`
	// UpdatedAt time.Time `json:"-"`
}
