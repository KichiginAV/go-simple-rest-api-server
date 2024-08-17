package repo

import (
	"database/sql"
	"errors"
	"simple-server/db"
	"simple-server/internal/model"

	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	db *sqlx.DB
}

func GetUsersRepo() UsersRepo {
	return &usersRepo{db.Get()}
}

func (u *usersRepo) GetUserByLogin(login string) (bool, *model.User, error) {
	user, err := u.sqlGetUserByLogin(login)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return false, nil, nil
		}
		return false, nil, err
	}

	return true, user, nil
}

func (u *usersRepo) CreateUser(user *model.User) error {
	err := u.sqlCreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
