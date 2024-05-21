package api

import (
	"github.com/gin-gonic/gin"
)

func (a *Application) SetupRoutes() {
	router := a.Server

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello! I'm alive!",
		})
	})

	//Cadastra entregador
	router.POST("/couriers", a.CreateCourier.Handler())

	//Cadastra localizações do entregador
	router.POST("/couriers/:courier_id/positions", a.UpdatePositions.Handler())

	//Lista localizações do entregador
	router.GET("/couriers/:courier_id/positions", a.GetPositions.Handler())

	//Lista última localização do entregador
	router.GET("/couriers/:courier_id/last-positions", a.GetLastPosition.Handler())

	//Lista de entregadores proximo a uma localização
	router.GET("/couriers/:lat/:lng")
}
