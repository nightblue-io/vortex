package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/nightblue-io/vortex/cmds"
	"github.com/nightblue-io/vortex/params"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	bold = color.New(color.Bold).SprintFunc()
	year = func() string {
		return fmt.Sprintf("%v", time.Now().Year())
	}

	rootCmd = &cobra.Command{
		Use:   "vortex",
		Short: bold("vortex") + " - Command line interface for Vortex",
		Long: bold("vortex") + ` - Command line interface for the Vortex Platform.
Copyright (c) ` + year() + ` NightBlue. All rights reserved.

The general form is ` + bold("vortex <resource[ subresource...]> <action> [flags]") + `. Most commands support
the ` + bold("--raw-input") + ` flag to be always in sync with the current feature set of the API in case the
built-in flags don't support all the possible input combinations yet. For beta APIs, we recommend
you to use the ` + bold("--raw-input") + ` flag.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			slog.Info("see -h for more information")
		},
	}
)

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.PersistentFlags().StringVar(&params.AuthProfile, "profile", params.AuthProfile, "profile name in ~/.config/nightblue/config.toml, default is [default]")
	rootCmd.PersistentFlags().StringVar(&params.ClientId, "client-id", os.Getenv("NIGHTBLUE_CLIENT_ID"), "your client id, defaults to $NIGHTBLUE_CLIENT_ID")
	rootCmd.PersistentFlags().StringVar(&params.ClientSecret, "client-secret", os.Getenv("NIGHTBLUE_CLIENT_SECRET"), "your client secret, defaults to $NIGHTBLUE_CLIENT_SECRET")
	rootCmd.PersistentFlags().StringVar(&params.OutFile, "out", params.OutFile, "output file, if the command supports writing to file")
	rootCmd.PersistentFlags().StringVar(&params.OutFmt, "outfmt", "csv", "output format: json, csv, valid if --out is set")
	rootCmd.PersistentFlags().BoolVar(&params.CleanOut, "bare", params.CleanOut, "if true, set console output to barebones, easier for scripting")
	rootCmd.AddCommand(
		cmds.WhoAmICmd(),
	)
}

func main() {
	cobra.EnableCommandSorting = false
	log.SetOutput(os.Stdout)
	rootCmd.Execute()
}
