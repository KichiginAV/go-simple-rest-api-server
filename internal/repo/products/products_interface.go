package repo

import "simple-server/internal/model"

type ProductsRepo interface {
	GetProductListOffset(offset int) (*[]model.Product, error)
	GetProductByID(id int) (bool, *model.Product, error)
	UpdateProduct(product *model.Product) error
	InsertProduct(product *model.Product) error
	DeleteProduct(id int) error
}
