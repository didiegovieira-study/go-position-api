package entity

import (
	"time"

	"github.com/google/uuid"
)

type Courier struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Email          string      `json:"email"`
	ActualPosition Positions   `json:"actual_position,omitempty"`
	Positions      []Positions `json:"positions,omitempty"`
}

type Positions struct {
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	CreatedAt string  `json:"created_at"`
}

func NewCourier(name, email string) *Courier {
	return &Courier{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Positions: []Positions{},
	}
}

func (c *Courier) AddLatLng(lat, lng float64) {
	c.ActualPosition = Positions{
		Lat:       lat,
		Lng:       lng,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	c.Positions = append(c.Positions, Positions{
		Lat:       lat,
		Lng:       lng,
		CreatedAt: time.Now().Format(time.RFC3339),
	})
}
