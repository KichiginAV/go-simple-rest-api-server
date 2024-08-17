package repo

import (
	"simple-server/db"
	"simple-server/internal/model"

	"github.com/jmoiron/sqlx"
)

type basketRepo struct {
	db *sqlx.DB
}

func GetBasketRepo() BasketRepo {
	return &basketRepo{db.Get()}
}

func (b *basketRepo) GetUserBasket(userID int) (*model.BasketUser, error) {
	uBasket, err := b.sqlGetUserBasket(userID)
	if err != nil {
		return nil, err
	}

	return uBasket, nil
}

func (b *basketRepo) AddUserBasket(basket *model.Basket) error {
	err := b.sqlAddUserBasket(basket)
	if err != nil {
		return err
	}

	return nil
}

func (b *basketRepo) DeleteProductOnBasket(basket *model.Basket) (int, error) {
	count, err := b.sqlDeleteProductOnBasket(basket)
	if err != nil {
		return count, err
	}

	return count, nil
}
