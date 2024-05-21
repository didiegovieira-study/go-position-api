package api

import (
	"github.com/didiegovieira/go-position-api/internal/infraestructure/api/handler"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Server *gin.Engine `wire:"-"`

	CreateCourier   *handler.CreateCourier
	GetPositions    *handler.GetPositions
	GetLastPosition *handler.GetLastPosition
	UpdatePositions *handler.UpdatePositions
}

func (a *Application) Start() error {
	a.Server = gin.Default()

	a.SetupRoutes()

	err := a.Server.Run(":3000")
	if err != nil {
		return err
	}

	return nil
}
