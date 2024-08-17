package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"simple-server/db"
	api "simple-server/internal/api/handlers"
	"simple-server/internal/auth"
	"simple-server/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var (
		confFile string
		conf     *config.Config
		err      error
	)

	flag.StringVar(&confFile, "c", "../../configs/config.yaml", "Configuration file path")
	flag.Parse()

	if conf, err = config.Init(confFile); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	db.InitDB(conf)
	defer db.CloseDB()

	e := echo.New()

	// Middleware для логирования
	// e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true, // Включаем поддержку учетных данных
	}))

	api.RegisterAPIRoute(e)
	auth.RegisterAuthRoute(e)

	e.Logger.Fatal(e.Start(conf.ListenHTTP))
}
