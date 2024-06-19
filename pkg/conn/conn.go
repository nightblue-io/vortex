package conn

import (
	"context"

	sdkcon "github.com/nightblue-io/vortex-go/conn"
)

func GetConnection(ctx context.Context, svcname string) (*sdkcon.GrpcClientConn, error) {
	return sdkcon.New(ctx, sdkcon.WithTargetService(svcname))
}
