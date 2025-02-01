package controllers

import (
	"github.com/gin-gonic/gin"
	"landingConstructor/app/services"
)

type PwaController interface {
	CreatePWA(c *gin.Context)
	CreatePreLanding(c *gin.Context)
	SaveImage(c *gin.Context)
	AddScreenshots(c *gin.Context)
}
type PwaControllerImpl struct {
	svc services.PwaService
}

func (p PwaControllerImpl) CreatePWA(c *gin.Context) {
	p.svc.CreatePWA(c)
}

func (p PwaControllerImpl) CreatePreLanding(c *gin.Context) {
	p.svc.CreatePreLanding(c)
}

func (p PwaControllerImpl) SaveImage(c *gin.Context) {
	p.svc.SaveImage(c)
}

func (p PwaControllerImpl) AddScreenshots(c *gin.Context) {
	p.svc.AddScreenshots(c)
}

func PwaControllerInit(pwaService services.PwaService) *PwaControllerImpl {
	return &PwaControllerImpl{
		svc: pwaService,
	}
}
