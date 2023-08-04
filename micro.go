package serviceGoMicro

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"go.uber.org/dig"
)

type GoMicroService struct {
	dig.In
	Server contracts.MicroService
}

func Serve() error {
	return container.Container.Invoke(func(svc GoMicroService) error {
		return svc.Server.Serve()
	})
}
