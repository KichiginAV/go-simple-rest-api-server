package repo

import (
	"database/sql"
	"errors"
	"simple-server/db"
	"simple-server/internal/model"

	"github.com/jmoiron/sqlx"
)

type productsRepo struct {
	db *sqlx.DB
}

func GetProductsRepo() ProductsRepo {
	return &productsRepo{db.Get()}
}

func (p *productsRepo) GetProductListOffset(offset int) (*[]model.Product, error) {
	products, err := p.sqlGetListOffset(offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productsRepo) GetProductByID(id int) (bool, *model.Product, error) {
	product, err := p.sqlGetByID(id)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, product, nil
}

func (p *productsRepo) UpdateProduct(product *model.Product) error {
	err := p.sqlUpdate(product)
	if err != nil {
		return err
	}

	return nil
}

func (p *productsRepo) InsertProduct(product *model.Product) error {
	err := p.sqlInsert(product)
	if err != nil {
		return err
	}

	return nil
}

func (p *productsRepo) DeleteProduct(id int) error {
	err := p.sqlDelete(id)
	if err != nil {
		return err
	}

	return nil
}
