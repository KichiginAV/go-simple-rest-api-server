package api

import (
	"log"
	"net/http"
	"simple-server/internal/handlers"
	"simple-server/internal/model"
	"simple-server/internal/usecases"
	"strconv"

	"github.com/labstack/echo/v4"
)

type basketHandler struct {
}

func NewBasketHandler() *basketHandler {
	return &basketHandler{}
}

func (b *basketHandler) GetUserBasket(c echo.Context) error {
	paramID := c.Param("userID")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid user ID"))
	}

	bUser, err := usecases.GetBasketUser(userID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErrf("error retrieving basket for user ID: %d", userID))
	}

	return c.JSON(http.StatusOK, bUser)
}

func (b *basketHandler) AddUserBasket(c echo.Context) error {
	paramID := c.Param("userID")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid user ID"))
	}

	basket := new(model.Basket)
	basket.UserID = userID
	if err := c.Bind(&basket); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid basket body"))
	}

	found, product, err := usecases.GetProductByID(basket.ProductID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErrf("error retrieving product by ID: %d", basket.ProductID))
	}

	if !found {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErrf("product with ID: %d not found", basket.ProductID))
	}

	basket.ProductID = product.ID

	err = usecases.AddBasketUser(basket)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErrf("error adding product by ID: %d to basket", basket.ProductID))
	}

	return c.JSON(http.StatusOK, handlers.ResponseOK("product added to basket"))
}

func (b *basketHandler) DeleteUserBasket(c echo.Context) error {
	paramID := c.Param("userID")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid user ID"))
	}

	basket := new(model.Basket)
	if err := c.Bind(&basket); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid basket body"))
	}
	basket.UserID = userID

	count, err := usecases.DeleteProductFromBasket(basket)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("error deleting product from basket"))
	}

	if count == 0 {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("no products deleted from basket"))
	}

	return c.JSON(http.StatusOK, handlers.ResponseOK("product deleted from basket"))
}
