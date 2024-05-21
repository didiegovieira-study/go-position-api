//go:build wireinject
// +build wireinject

package di

import (
	"github.com/didiegovieira/go-position-api/internal/infraestructure/api"
	"github.com/google/wire"
)

var wireApiSet = wire.NewSet(
	provideMongoDbClient,

	provideCourierRepositoryMongodb,

	provideCreateCourierUseCaseSet,
	provideUpdatePositionsUseCaseSet,
	provideGetPositionsUseCaseSet,
	provideGetLastPositionUseCaseSet,

	apiHandlersSet,
	wire.Struct(new(api.Application), "*"),
)

func InitializeApi() (*api.Application, func(), error) {
	wire.Build(wireApiSet)
	return &api.Application{}, func() {}, nil
}
