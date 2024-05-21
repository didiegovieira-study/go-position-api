package di

import (
	"github.com/didiegovieira/go-position-api/internal/application/usecase"
	"github.com/google/wire"
)

var provideCreateCourierUseCaseSet = wire.NewSet(
	usecase.NewCreateCourier,
	wire.Bind(new(usecase.CreateCourierInterface), new(*usecase.CreateCourier)),
)

var provideUpdatePositionsUseCaseSet = wire.NewSet(
	usecase.NewUpdatePositions,
	wire.Bind(new(usecase.UpdatePositionsInterface), new(*usecase.UpdatePositions)),
)

var provideGetPositionsUseCaseSet = wire.NewSet(
	usecase.NewGetPositions,
	wire.Bind(new(usecase.GetPositionsInterface), new(*usecase.GetPositions)),
)

var provideGetLastPositionUseCaseSet = wire.NewSet(
	usecase.NewGetLastPosition,
	wire.Bind(new(usecase.GetLastPositionInterface), new(*usecase.GetLastPosition)),
)
