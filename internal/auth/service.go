package auth

import (
	"log"
	"net/http"
	"simple-server/internal/handlers"
	"simple-server/internal/model"
	"simple-server/internal/usecases"
	"simple-server/pkg/validators"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (u *userHandler) RegisterUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErrf("invalid user body"))
	}

	found, _, err := usecases.GetUserByLogin(user.Login)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErr("error retrieving user"))
	}

	if found {
		return c.JSON(http.StatusConflict, handlers.ResponseErr("user already exists"))
	}

	if err := validators.ValidateLogin(user.Login); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErrf("invalid user login: %v", err))
	}

	if err := validators.ValidatePassword(user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ResponseErrf("invalid user password: %v", err))
	}

	hash, err := hashPassword(user.Password) // Changed from hashPassword to validators.HashPassword
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErr("error creating hash password"))
	}
	user.PasswordHash = hash

	err = usecases.RegisterUser(user) // Changed from RegistrationUser to RegisterUser
	if err != nil {
		return c.JSON(http.StatusInternalServerError, handlers.ResponseErr("error saving user"))
	}

	return c.JSON(http.StatusOK, handlers.ResponseOK("user created successfully"))
}
