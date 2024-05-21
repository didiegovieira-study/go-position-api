package usecase

import (
	"github.com/didiegovieira/go-position-api/internal/application/dto"
	"github.com/didiegovieira/go-position-api/internal/domain/entity"
)

type UpdatePositionsInterface interface {
	Execute(input dto.UpdatePositionsInput) ([]entity.Positions, error)
}
