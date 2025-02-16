package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"landingConstructor/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()

	router.Use(cors.Default())

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Static("/uploads", "./uploads")

	api := router.Group("/api/v1")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.GET("/common", init.CommonCtrl.Ping)
	{
		pwa := api.Group("/pwa")
		pwa.POST("", init.PwaCtrl.CreatePWA)
		pwa.POST("/prelanding", init.PwaCtrl.CreatePreLanding)
		pwa.POST("/logo", init.PwaCtrl.SaveImage)
		pwa.POST("/screenshots", init.PwaCtrl.AddScreenshots)
	}

	return router
}
