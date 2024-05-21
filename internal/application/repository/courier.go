package repository

import "github.com/didiegovieira/go-position-api/internal/domain/entity"

type Courier interface {
	Save(courier *entity.Courier) error
	GetById(id string) (*entity.Courier, error)
}
