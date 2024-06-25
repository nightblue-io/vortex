package login

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/nightblue-io/vortex-go/iam/v1"
	"github.com/nightblue-io/vortex/internal"
	"github.com/nightblue-io/vortex/internal/conn"
	"github.com/nightblue-io/vortex/params"
)

func RefreshToken() error {
	ctx := context.Background()
	gcon, err := conn.GetConnection(ctx, &conn.GetConnectionOptions{
		Target:      params.Addr,
		ServiceName: "iam",
	})

	if err != nil {
		return err
	}

	client, err := iam.NewClient(ctx, &iam.ClientOptions{Conn: gcon})
	if err != nil {
		return err
	}

	defer client.Close()
	stream, err := client.Login(ctx, &iam.LoginRequest{
		AccessToken: internal.GetLocalAccessToken(),
	})

	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch {
		case resp.AccessToken != "":
			os.MkdirAll(internal.DirVortex, os.ModePerm)
			os.WriteFile(internal.FileAccessToken, []byte(resp.AccessToken), 0600)
		}
	}

	return nil
}

func DeviceAuth() error {
	ctx := context.Background()
	gcon, err := conn.GetConnection(ctx, &conn.GetConnectionOptions{
		Target:      params.Addr,
		ServiceName: "iam",
	})

	if err != nil {
		return err
	}

	client, err := iam.NewClient(ctx, &iam.ClientOptions{Conn: gcon})
	if err != nil {
		return err
	}

	defer client.Close()
	stream, err := client.Login(ctx, &iam.LoginRequest{DeviceAuth: "1"})
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			slog.Error("failed:", "error", err)
			return err
		}

		switch {
		case resp.DeviceAuth != nil:
			var m strings.Builder
			fmt.Fprintf(&m, "Your user code is %v. ", resp.DeviceAuth.UserCode)
			fmt.Fprintf(&m, "Open the URL below in your browser to complete the login.\n\n")
			fmt.Fprintf(&m, "\t%v\n\n", resp.DeviceAuth.VerificationUriComplete)
			tm := time.Second * time.Duration(resp.DeviceAuth.ExpiresIn)
			fmt.Fprintf(&m, "Your user code (and URL) will expire in %v.\n", tm)
			fmt.Fprintf(&m, "Open the link? [Y/n]: ")
			fmt.Printf(m.String())
			input := "Y"
			fmt.Scanln(&input)
			if strings.ToUpper(input) == "Y" {
				go internal.OpenUrl(resp.DeviceAuth.VerificationUriComplete)
			}
		case resp.AccessToken != "":
			os.MkdirAll(internal.DirVortex, os.ModePerm)
			os.WriteFile(internal.FileAccessToken, []byte(resp.AccessToken), 0600)

			var m strings.Builder
			fmt.Fprintf(&m, "Your access token is:\n\n")
			fmt.Fprintf(&m, "\t%v\n\n", resp.AccessToken)
			fmt.Fprintf(&m, "It's been saved to %v for later calls.", internal.FileAccessToken)
			fmt.Println(m.String())
		}
	}

	return nil
}
