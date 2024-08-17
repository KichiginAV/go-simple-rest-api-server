package repo

import (
	"fmt"
	"simple-server/internal/model"
)

var (
	sqlGetUserBasket = `SELECT p.id, p.label, p.description, p.price
	FROM basket
	JOIN products p ON p.id = basket.product_id
	WHERE basket.user_id = $1`
	sqlAddUserBasket = `INSERT INTO basket(user_id,product_id)
	VALUES($1,$2)`
	sqlDeleteProductOnBasket = `WITH cte AS (
    SELECT ctid
    FROM basket
    WHERE user_id = $1 AND product_id = $2
    LIMIT 1
	)
	DELETE FROM basket
	WHERE ctid IN (SELECT ctid FROM cte);`
)

func (b *basketRepo) sqlGetUserBasket(userID int) (*model.BasketUser, error) {
	bUser := &model.BasketUser{UserID: userID}

	err := b.db.Select(&bUser.Products, sqlGetUserBasket, userID)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}

	return bUser, nil
}

func (b *basketRepo) sqlAddUserBasket(basket *model.Basket) error {
	_, err := b.db.Exec(sqlAddUserBasket, basket.UserID, basket.ProductID)
	if err != nil {
		return err
	}

	return nil
}

func (b *basketRepo) sqlDeleteProductOnBasket(basket *model.Basket) (int, error) {
	res, err := b.db.Exec(sqlDeleteProductOnBasket, basket.UserID, basket.ProductID)
	rowCount, _ := res.RowsAffected()
	if err != nil {
		return int(rowCount), err
	}

	return int(rowCount), nil
}
