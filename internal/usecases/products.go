package usecases

import (
	"fmt"
	"simple-server/internal/model"
	repo "simple-server/internal/repo/products"
)

func GetProductByID(id int) (bool, *model.Product, error) {
	found, product, err := repo.GetProductsRepo().GetProductByID(id)
	if err != nil {
		return false, nil, fmt.Errorf("getting product by id: %v", err)
	}

	return found, product, nil
}

func GetProductsListOffset(offset int) (*[]model.Product, error) {
	products, err := repo.GetProductsRepo().GetProductListOffset(offset)
	if err != nil {
		return nil, fmt.Errorf("getting product by id: %v", err)
	}

	return products, nil
}

func UpdateProduct(product *model.Product) error {
	err := repo.GetProductsRepo().UpdateProduct(product)
	if err != nil {
		return fmt.Errorf("update product: %v", err)
	}

	return nil
}

func InsertProduct(product *model.Product) error {
	err := repo.GetProductsRepo().InsertProduct(product)
	if err != nil {
		return fmt.Errorf("insert product: %v", err)
	}

	return nil
}

func DeleteProduct(id int) error {
	err := repo.GetProductsRepo().DeleteProduct(id)
	if err != nil {
		return fmt.Errorf("delete product with id: %d - %v", id, err)
	}

	return nil
}
