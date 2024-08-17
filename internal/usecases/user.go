package usecases

import (
	"fmt"
	"simple-server/internal/model"
	repo "simple-server/internal/repo/users"
)

func RegisterUser(user *model.User) error {
	err := repo.GetUsersRepo().CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByLogin(login string) (bool, *model.User, error) {
	found, user, err := repo.GetUsersRepo().GetUserByLogin(login)
	if err != nil {
		return false, nil, fmt.Errorf("get user by login: %v", err)
	}

	return found, user, nil
}
