package router

import (
	"healthcheck/config"
	"healthcheck/controller"
	"healthcheck/usecase"

	"github.com/gin-gonic/gin"
)

//New -
func New(config config.Interface) *gin.Engine {

	usecase := usecase.New(config)
	ctrl := controller.New(usecase)

	gin.SetMode(config.GinMode())

	app := gin.Default()

	baseURL := config.AppURL()
	router := app.Group(baseURL)

	router.Use(gin.Recovery())

	router.GET("/line/authorization", ctrl.GetLineAuthen)
	router.GET("/healthcheck", ctrl.GetHealthCheckStatus)
	router.POST("/line/authorization", ctrl.PostLineAuthen)

	return app
}
