package usecase

import (
	"github.com/didiegovieira/go-position-api/internal/application/repository"
	"github.com/didiegovieira/go-position-api/internal/domain/entity"
)

type GetPositions struct {
	CourierRepository repository.Courier
}

func NewGetPositions(courierRepository repository.Courier) *GetPositions {
	return &GetPositions{
		CourierRepository: courierRepository,
	}
}

func (g *GetPositions) Execute(courierID string) ([]entity.Positions, error) {
	courier, err := g.CourierRepository.GetById(courierID)
	if err != nil {
		return nil, err
	}

	return courier.Positions, nil
}
