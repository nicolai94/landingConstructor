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

// HelloHandler обрабатывает запрос на /hello
// @Summary Возвращает приветствие
// @Description Возвращает простое приветствие
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} string "Привет, мир!"
// @Router /hello [get]
func (u CommonServiceImpl) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	return
}

func CommonServiceInit() CommonService {
	return &CommonServiceImpl{}
}
