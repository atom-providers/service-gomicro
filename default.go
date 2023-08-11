package serviceGoMicro

import (
	"github.com/atom-providers/app"
	"github.com/atom-providers/log"
	microGoMicro "github.com/atom-providers/micro-gomicro"
	"github.com/atom-providers/uuid"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		app.DefaultProvider(),
		log.DefaultProvider(),
		uuid.DefaultProvider(),
		microGoMicro.DefaultProvider(),
	}, providers...)
}
