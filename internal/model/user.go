package model

type User struct {
	ID           int    `json:"id" db:"id"`
	Login        string `json:"login" db:"login"`
	Name         string `json:"name" db:"name"`
	PasswordHash string `json:"passhash" db:"passhash"`
	Password     string `json:"password" db:"-"`
}
