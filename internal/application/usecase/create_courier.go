package usecase

import (
	"github.com/didiegovieira/go-position-api/internal/application/dto"
	"github.com/didiegovieira/go-position-api/internal/application/repository"
	"github.com/didiegovieira/go-position-api/internal/domain/entity"
)

type CreateCourier struct {
	CourierRepository repository.Courier
}

func NewCreateCourier(courierRepository repository.Courier) *CreateCourier {
	return &CreateCourier{
		CourierRepository: courierRepository,
	}
}

func (c *CreateCourier) Execute(input dto.CreateCourierInput) (*dto.CreateCourierOutput, error) {
	courier := entity.NewCourier(input.Name, input.Email)

	err := c.CourierRepository.Save(courier)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCourierOutput{
		ID:    courier.ID,
		Name:  courier.Name,
		Email: courier.Email,
	}, nil
}
