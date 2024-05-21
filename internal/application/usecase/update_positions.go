package usecase

import (
	"github.com/didiegovieira/go-position-api/internal/application/dto"
	"github.com/didiegovieira/go-position-api/internal/application/repository"
	"github.com/didiegovieira/go-position-api/internal/domain/entity"
)

type UpdatePositions struct {
	CourierRepository repository.Courier
}

func NewUpdatePositions(courierRepository repository.Courier) *UpdatePositions {
	return &UpdatePositions{
		CourierRepository: courierRepository,
	}
}

func (u *UpdatePositions) Execute(input dto.UpdatePositionsInput) ([]entity.Positions, error) {
	courier, err := u.CourierRepository.GetById(input.ID)
	if err != nil {
		return nil, err
	}

	courier.AddLatLng(input.Lat, input.Lng)

	err = u.CourierRepository.Save(courier)
	if err != nil {
		return nil, err
	}

	return courier.Positions, nil
}
