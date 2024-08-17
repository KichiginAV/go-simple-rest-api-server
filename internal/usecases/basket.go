package usecases

import (
	"fmt"
	"simple-server/internal/model"
	repo "simple-server/internal/repo/basket"
)

func GetBasketUser(id int) (*model.BasketUser, error) {
	bUser, err := repo.GetBasketRepo().GetUserBasket(id)
	if err != nil {
		return nil, fmt.Errorf("getting basket by user id: %d - %v", id, err)
	}

	return bUser, nil
}

func AddBasketUser(basket *model.Basket) error {
	err := repo.GetBasketRepo().AddUserBasket(basket)
	if err != nil {
		return fmt.Errorf("add basket with user id: %d and product id: %d - %v", basket.UserID, basket.ProductID, err)
	}

	return nil
}

func DeleteProductFromBasket(basket *model.Basket) (int, error) {
	count, err := repo.GetBasketRepo().DeleteProductOnBasket(basket)
	if err != nil {
		return count, fmt.Errorf("delete basket with user id: %d and product id: %d - %v", basket.UserID, basket.ProductID, err)
	}

	return count, nil
}
