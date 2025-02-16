package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonService interface {
	Ping(c *gin.Context)
}

type CommonServiceImpl struct {
}

// PingHamdler обрабатывает запрос на ping
// @Summary Router for ping
// @Description Pinging for server and app
// @Tags Common
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /common [get]
func (u CommonServiceImpl) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return
}

func CommonServiceInit() CommonService {
	return &CommonServiceImpl{}
}
