package repo

import "simple-server/internal/model"

var (
	sqlGetUserByLogin = `SELECT * FROM users
	WHERE login = $1`
	sqlCreateUser = `INSERT INTO users(login,name,passhash)
	VALUES($1,$2,$3)`
)

func (u *usersRepo) sqlGetUserByLogin(login string) (*model.User, error) {
	user := new(model.User)
	err := u.db.Get(user, sqlGetUserByLogin, login)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *usersRepo) sqlCreateUser(user *model.User) error {
	_, err := u.db.Exec(sqlCreateUser, user.Login, user.Name, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}
