package handler

import (
	"github.com/didiegovieira/go-position-api/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type GetLastPosition struct {
	GetLastPositionUseCase usecase.GetLastPositionInterface
}

func (g *GetLastPosition) Handler() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		courierID := ctx.Param("courier_id")

		output, err := g.GetLastPositionUseCase.Execute(courierID)
		if err != nil {
			ctx.JSON(400, gin.H{})
			return
		}

		ctx.JSON(200, gin.H{
			"data": output,
		})
	}
}
