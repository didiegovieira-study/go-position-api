package handler

import (
	"github.com/didiegovieira/go-position-api/internal/application/dto"
	"github.com/didiegovieira/go-position-api/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type CreateCourier struct {
	CreateCourierUseCase usecase.CreateCourierInterface
}

func (c *CreateCourier) Handler() func(c *gin.Context) {
	return func(ctx *gin.Context) {
		var input dto.CreateCourierInput

		err := ctx.BindJSON(&input)
		if err != nil {
			ctx.JSON(400, gin.H{})
			return
		}

		output, err := c.CreateCourierUseCase.Execute(input)
		if err != nil {
			ctx.JSON(400, gin.H{})
			return
		}

		ctx.JSON(200, gin.H{
			"data": output,
		})
	}
}
