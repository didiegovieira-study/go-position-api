package di

import (
	repository2 "github.com/didiegovieira/go-position-api/internal/application/repository"
	"github.com/didiegovieira/go-position-api/internal/infraestructure/repository"
	"github.com/google/wire"
)

var provideCourierRepositoryMongodb = wire.NewSet(
	repository.NewCourierMongodb,
	wire.Bind(new(repository2.Courier), new(*repository.CourierMongodb)),
)
