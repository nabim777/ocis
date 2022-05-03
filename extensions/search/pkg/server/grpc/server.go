package grpc

import (
	svc "github.com/owncloud/ocis/extensions/search/pkg/service/v0"
	"github.com/owncloud/ocis/ocis-pkg/service/grpc"
	"github.com/owncloud/ocis/ocis-pkg/version"
	searchsvc "github.com/owncloud/ocis/protogen/gen/ocis/services/search/v0"
)

// Server initializes a new go-micro service ready to run
func Server(opts ...Option) grpc.Service {
	options := newOptions(opts...)

	service := grpc.NewService(
		grpc.Name(options.Config.Service.Name),
		grpc.Context(options.Context),
		grpc.Address(options.Config.GRPC.Addr),
		grpc.Namespace(options.Config.GRPC.Namespace),
		grpc.Logger(options.Logger),
		grpc.Flags(options.Flags...),
		grpc.Version(version.String),
	)

	handle, err := svc.NewHandler(
		svc.Config(options.Config),
		svc.Logger(options.Logger),
	)
	if err != nil {
		options.Logger.Error().
			Err(err).
			Msg("Error initializing search service")
		return grpc.Service{}
	}
	_ = searchsvc.RegisterSearchProviderHandler(
		service.Server(),
		handle,
	)
	return service
}