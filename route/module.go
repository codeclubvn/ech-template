package route

import (
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Invoke(NewUserRoutes))
