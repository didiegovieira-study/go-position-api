package handler

import (
	"github.com/didiegovieira/go-position-api/internal/application/dto"
	"github.com/didiegovieira/go-position-api/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type UpdatePositions struct {
	UpdatePositionsUseCase usecase.UpdatePositionsInterface
}

func (u *UpdatePositions) Handler() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		var input dto.UpdatePositionsInput

		err := ctx.BindJSON(&input)
		if err != nil {
			ctx.JSON(400, gin.H{})
			return
		}

		input.ID = ctx.Param("courier_id")

		output, err := u.UpdatePositionsUseCase.Execute(input)
		if err != nil {
			ctx.JSON(400, gin.H{})
			return
		}

		ctx.JSON(200, gin.H{
			"data": output,
		})
	}
}
