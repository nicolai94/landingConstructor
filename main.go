package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	_ "github.com/swaggo/gin-swagger/landingConstructor/docs"
	"landingConstructor/app/routers"
	"landingConstructor/config"
	"os"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server for a Gin Swagger example.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Ошибка загрузки .env файла:", err)
	}
	init := config.InitDependencies()
	db := config.ConnectToDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := goose.Up(sqlDB, "migrations"); err != nil {
		panic(err)
	}
	app := routers.Init(init)

	if err := app.Run(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
