package api

import (
	"net/http"
	"simple-server/internal/handlers"
	"simple-server/internal/model"
	"simple-server/internal/usecases"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
}

func NewProductHandler() *productHandler {
	return &productHandler{}
}

func (p *productHandler) GetList(c echo.Context) error {
	queryParam := c.QueryParam("offset")
	offset, err := strconv.Atoi(queryParam)
	if err != nil || queryParam == "" {
		offset = 0
	}

	products, err := usecases.GetProductsListOffset(offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErr("error retrieving products list"))
	}

	return c.JSON(http.StatusOK, products)
}

func (p *productHandler) GetByID(c echo.Context) error {
	paramID := c.Param("id")
	productID, err := strconv.Atoi(paramID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid product ID"))
	}

	found, product, err := usecases.GetProductByID(productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErr("error retrieving product"))
	}

	if !found {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErrf("product with ID: %d not found", productID))
	}

	return c.JSON(http.StatusOK, product)
}

func (p *productHandler) Update(c echo.Context) error {
	product := new(model.Product)
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid product body"))
	}

	if product.Price == 0 || product.Label == "" || product.Description == "" {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid product body"))
	}

	found, oldProduct, err := usecases.GetProductByID(product.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErr("error updating product"))
	}

	if !found {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErrf("product with ID: %d not found", product.ID))
	}

	if product.Description != oldProduct.Description {
		oldProduct.Description = product.Description
	}

	if product.Label != oldProduct.Label {
		oldProduct.Label = product.Label
	}

	if product.Price != oldProduct.Price {
		oldProduct.Price = product.Price
	}

	err = usecases.UpdateProduct(oldProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErr("error updating product"))
	}

	return c.JSON(http.StatusOK, handlers.ResponseOKf("product with ID: %d was updated", product.ID))
}

func (p *productHandler) Insert(c echo.Context) error {
	product := new(model.Product)
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid product body"))
	}

	if product.Price == 0 && product.Label == "" && product.Description == "" {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid product body"))
	}

	err := usecases.InsertProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErr("error inserting product"))
	}

	return c.JSON(http.StatusOK, handlers.ResponseOKf("product with ID: %d was inserted", product.ID))
}

func (p *productHandler) Delete(c echo.Context) error {
	queryParam := c.QueryParam("id")
	id, err := strconv.Atoi(queryParam)
	if err != nil || queryParam == "" {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErr("invalid product ID"))
	}

	err = usecases.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErrf("error deleting product with ID: %s", queryParam))
	}

	return c.JSON(http.StatusOK, handlers.ResponseOK("product was deleted"))
}
