package api

import (
	"github.com/labstack/echo/v4"
)

func RegisterAPIRoute(e *echo.Echo) {
	v1 := e.Group("api/v1")

	products := NewProductHandler()
	productsGroup := v1.Group("/products")
	productsGroup.GET("", products.GetList)
	productsGroup.GET("/:id", products.GetByID)
	productsGroup.PUT("", products.Update)
	productsGroup.POST("", products.Insert)
	productsGroup.DELETE("/:id", products.Delete)

	basket := NewBasketHandler()
	basketGroup := v1.Group("/basket")
	basketGroup.GET("/:userID", basket.GetUserBasket)
	basketGroup.POST("/:userID", basket.AddUserBasket)
	basketGroup.DELETE("/:userID", basket.DeleteUserBasket)

}
