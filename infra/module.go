package infra

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(NewDatabase, NewCloudinary))
