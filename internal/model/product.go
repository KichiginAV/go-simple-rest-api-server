package model

import "encoding/json"

type Product struct {
	ID          int    `json:"id" db:"id"`
	Label       string `json:"label" db:"label"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
}

type ProductsList []Product

func (p ProductsList) MarshalJSON() ([]byte, error) {
	if p == nil {
		return json.Marshal(make([]Product, 0))
	}

	return json.Marshal([]Product(p))
}
