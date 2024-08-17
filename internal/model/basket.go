package model

type Basket struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

type BasketUser struct {
	UserID   int          `json:"user_id" db:"user_id"`
	Products ProductsList `json:"products" db:"products"`
}
