package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"landingConstructor/app/routers"
	"landingConstructor/config"
	_ "landingConstructor/docs"
	"os"
)

// @title Gin Swagger Landing Constructor
// @version 1.0
// @description App for landing constructor.

// @host localhost:8080
// @BasePath /api
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Ошибка загрузки .env файла:", err)
	}
	init := config.InitDependencies()
	db := config.ConnectToDB()

	sqlDB, dbErr := db.DB()
	if dbErr != nil {
		panic(dbErr)
	}
	if migrationErr := goose.Up(sqlDB, "migrations"); migrationErr != nil {
		panic(migrationErr)
	}
	app := routers.Init(init)

	if err := app.Run(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
