//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"landingConstructor/app/controllers"
	"landingConstructor/app/repositories"
	"landingConstructor/app/services"
)

func InitDependencies() *Initialization {
	db := ConnectToDB()

	commonService := services.CommonServiceInit()
	commonControllerImpl := controllers.CommonControllerInit(commonService)

	pwaRepoInit := repositories.PwaRepositoryInit(db)              // Репозиторий инициализируется с базой данных
	pwaService := services.PwaServiceInit(pwaRepoInit)             // Передаем репозиторий в сервис
	pwaControllerImpl := controllers.PwaControllerInit(pwaService) // Передаем сервис в контроллер

	initialization := NewInitialization(
		commonService,
		commonControllerImpl,
		pwaService,
		pwaControllerImpl,
		pwaRepoInit,
	)
	return initialization
}
