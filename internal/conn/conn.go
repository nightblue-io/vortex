package conn

import (
	"context"

	sdkcon "github.com/nightblue-io/vortex-go/conn"
)

type GetConnectionOptions struct {
	Target      string
	ServiceName string
}

func GetConnection(ctx context.Context, opt *GetConnectionOptions) (*sdkcon.GrpcClientConn, error) {
	opts := []sdkcon.ClientOption{}
	if opt.Target != "" {
		opts = append(opts, sdkcon.WithTarget(opt.Target))
	}

	if opt.ServiceName != "" {
		opts = append(opts, sdkcon.WithTargetService(opt.ServiceName))
	}

	return sdkcon.New(ctx, opts...)
}
