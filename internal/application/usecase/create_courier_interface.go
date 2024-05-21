package usecase

import "github.com/didiegovieira/go-position-api/internal/application/dto"

type CreateCourierInterface interface {
	Execute(input dto.CreateCourierInput) (*dto.CreateCourierOutput, error)
}
