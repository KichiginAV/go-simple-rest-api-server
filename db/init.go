package db

import (
	"fmt"
	"log"
	"os"
	"simple-server/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const connStr = "%s://%s:%s@%s:%d/%s?sslmode=disable"

func InitDB(conf *config.Config) {
	connStr := fmt.Sprintf(connStr,
		conf.DataBase.Driver,
		conf.DataBase.User,
		conf.DataBase.Password,
		conf.DataBase.Address,
		conf.DataBase.Port,
		conf.DataBase.DBName)

	var err error
	db, err = sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	if err != nil {
		log.Fatalf("Unable to create connection ping: %v\n", err)
		os.Exit(1)
	}
	log.Println("Connected to database")
}

func Get() *sqlx.DB {
	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
