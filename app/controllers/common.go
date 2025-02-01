package controllers

import (
	"github.com/gin-gonic/gin"
	"landingConstructor/app/services"
)

type CommonController interface {
	Ping(c *gin.Context)
}
type CommonControllerImpl struct {
	svc services.CommonService
}

func (u CommonControllerImpl) Ping(c *gin.Context) {
	u.svc.Ping(c)
}

func CommonControllerInit(commonService services.CommonService) *CommonControllerImpl {
	return &CommonControllerImpl{
		svc: commonService,
	}
}
