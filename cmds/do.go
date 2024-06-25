package cmds

import (
	"context"
	"log/slog"
	"os"

	"github.com/nightblue-io/vortex-go/vortex/v1"
	"github.com/nightblue-io/vortex/internal/conn"
	"github.com/nightblue-io/vortex/params"
	"github.com/spf13/cobra"
)

func DoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "do",
		Short: "Test cmd for server",
		Long:  `Test command for server.`,
		Run: func(cmd *cobra.Command, args []string) {
			var ret int
			defer func(r *int) {
				if *r != 0 {
					os.Exit(*r)
				}
			}(&ret)

			fnerr := func(e error) {
				slog.Error("failed:", "err", e)
				ret = 1
			}

			ctx := context.Background()
			gcon, err := conn.GetConnection(ctx, &conn.GetConnectionOptions{
				Target:      params.Addr,
				ServiceName: "vortex",
			})

			if err != nil {
				fnerr(err)
				return
			}

			client, err := vortex.NewClient(ctx, &vortex.ClientOptions{Conn: gcon})
			if err != nil {
				fnerr(err)
				return
			}

			defer client.Close()
			resp, err := client.Do(ctx, &vortex.DoRequest{})
			if err != nil {
				fnerr(err)
				return
			}

			slog.Info("dbg:", "data", resp.Data)
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}
