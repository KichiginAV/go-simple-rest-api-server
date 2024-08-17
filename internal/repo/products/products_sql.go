package repo

import (
	"simple-server/internal/model"
)

var (
	sqlSelectOffset = `SELECT * FROM products
	ORDER by id
	LIMIT 10
	OFFSET $1`
	sqlGetByID = `SELECT * FROM products
	WHERE id = $1`
	sqlUpdate = `UPDATE products
	SET label = $1,
	description = $2,
	price = $3
	WHERE id = $4`
	sqlInsert = `INSERT INTO products(label,description,price)
	VALUES($1,$2,$3)`
	sqlDelete = `DELETE FROM products
	WHERE id = $1`
)

func (p *productsRepo) sqlGetListOffset(offset int) (*[]model.Product, error) {
	products := new([]model.Product)
	err := p.db.Select(&products, sqlSelectOffset, offset)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productsRepo) sqlGetByID(id int) (*model.Product, error) {
	product := new(model.Product)
	err := p.db.Get(product, sqlGetByID, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *productsRepo) sqlUpdate(product *model.Product) error {
	_, err := p.db.Exec(sqlUpdate, product.Label, product.Description, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *productsRepo) sqlInsert(product *model.Product) error {
	_, err := p.db.Exec(sqlInsert, product.Label, product.Description, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func (p *productsRepo) sqlDelete(id int) error {
	_, err := p.db.Exec(sqlDelete, id)
	if err != nil {
		return err
	}

	return nil
}
