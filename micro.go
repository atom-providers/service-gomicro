package serviceGoMicro

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/contracts"
	"github.com/spf13/cobra"
	"go-micro.dev/v4"
	"go.uber.org/dig"
)

type GoMicroService struct {
	dig.In
	Server   contracts.MicroService
	Initials []contracts.Initial `group:"initials"`
}

func Serve() error {
	return container.Container.Invoke(func(svc GoMicroService) error {
		service := svc.Server.GetEngine().(micro.Service)
		service.Init()

		return svc.Server.Serve()
	})
}

func ServeE(cmd *cobra.Command, args []string) error {
	return Serve()
}
