package usecase

import (
	"github.com/didiegovieira/go-position-api/internal/application/repository"
	"github.com/didiegovieira/go-position-api/internal/domain/entity"
)

type GetLastPosition struct {
	CourierRepository repository.Courier
}

func NewGetLastPosition(courierRepository repository.Courier) *GetLastPosition {
	return &GetLastPosition{
		CourierRepository: courierRepository,
	}
}

func (g *GetLastPosition) Execute(courierID string) (*entity.Positions, error) {
	courier, err := g.CourierRepository.GetById(courierID)
	if err != nil {
		return nil, err
	}

	return &courier.Positions[len(courier.Positions)-1], nil
}
