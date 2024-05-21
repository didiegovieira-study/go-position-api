package usecase

import "github.com/didiegovieira/go-position-api/internal/domain/entity"

type GetLastPositionInterface interface {
	Execute(courierID string) (*entity.Positions, error)
}
