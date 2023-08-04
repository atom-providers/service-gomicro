package serviceGoMicro

import (
	"github.com/atom-providers/log"
	microGoMicro "github.com/atom-providers/micro-gomicro"
	"github.com/rogeecn/atom/container"
)

func Default(providers ...container.ProviderContainer) container.Providers {
	return append(container.Providers{
		log.DefaultProvider(),
		microGoMicro.DefaultProvider(),
	}, providers...)
}
