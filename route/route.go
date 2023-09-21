package route

import (
	"go.uber.org/fx"
	"trail_backend/route/v1"
)

var Module = fx.Options(fx.Invoke(v1.NewAuthRoutes, v1.NewUserRoutes, v1.NewPostRoutes, v1.NewFileRoutes))