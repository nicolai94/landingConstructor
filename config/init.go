package config

import (
	"landingConstructor/app/controllers"
	"landingConstructor/app/repositories"
	"landingConstructor/app/services"
)

type Initialization struct {
	commonSvc  services.CommonService
	CommonCtrl controllers.CommonController
	pwaSvc     services.PwaService
	PwaCtrl    controllers.PwaController
	PwaRepo    repositories.PwaRepository
}

func NewInitialization(
	commonService services.CommonService,
	commonCtrl controllers.CommonController,
	pwaService services.PwaService,
	pwaCtrl controllers.PwaController,
	pwaRepo repositories.PwaRepository,
) *Initialization {
	return &Initialization{
		commonSvc:  commonService,
		CommonCtrl: commonCtrl,
		pwaSvc:     pwaService,
		PwaCtrl:    pwaCtrl,
		PwaRepo:    pwaRepo,
	}
}
