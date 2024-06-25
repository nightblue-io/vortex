package cmds

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nightblue-io/vortex/internal"
	"github.com/nightblue-io/vortex/internal/login"
	"github.com/spf13/cobra"
)

func LoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to Vortex",
		Long:  `Login to Vortex.`,
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

			token := internal.GetLocalAccessToken()
			switch {
			case token != "":
				err := login.RefreshToken()
				if err != nil {
					fnerr(err)
					return
				}

				fmt.Println("Successful. Your access token has been updated.")
			default:
				err := login.DeviceAuth()
				if err != nil {
					fnerr(err)
					return
				}
			}
		},
	}

	cmd.Flags().SortFlags = false
	return cmd
}
