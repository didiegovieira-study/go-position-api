package di

import (
	"github.com/didiegovieira/go-position-api/internal/infraestructure/api/handler"
	"github.com/google/wire"
)

var apiHandlersSet = wire.NewSet(
	wire.Struct(new(handler.CreateCourier), "*"),
	wire.Struct(new(handler.GetPositions), "*"),
	wire.Struct(new(handler.GetLastPosition), "*"),
	wire.Struct(new(handler.UpdatePositions), "*"),
)
