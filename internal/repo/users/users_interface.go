package repo

import "simple-server/internal/model"

type UsersRepo interface {
	GetUserByLogin(login string) (bool, *model.User, error)
	CreateUser(user *model.User) error
}
