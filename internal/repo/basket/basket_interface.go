package repo

import "simple-server/internal/model"

type BasketRepo interface {
	GetUserBasket(userID int) (*model.BasketUser, error)
	AddUserBasket(basket *model.Basket) error
	DeleteProductOnBasket(basket *model.Basket) (int, error)
}
