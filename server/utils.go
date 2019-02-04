package server

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"

	"github.com/ironman0x7b2/sentinel-sdk/version"
)

func AddCommands(ctx *server.Context, cdc *codec.Codec,
	rootCmd *cobra.Command,
	appCreator server.AppCreator, appExport server.AppExporter) {

	rootCmd.PersistentFlags().String("log_level", ctx.Config.LogLevel, "Log level")

	tendermintCmd := &cobra.Command{
		Use:   "tendermint",
		Short: "Tendermint subcommands",
	}

	tendermintCmd.AddCommand(
		server.ShowNodeIDCmd(ctx),
		server.ShowValidatorCmd(ctx),
		server.ShowAddressCmd(ctx),
		server.VersionCmd(ctx),
	)

	rootCmd.AddCommand(
		server.StartCmd(ctx, appCreator),
		server.UnsafeResetAllCmd(ctx),
		client.LineBreak,
		tendermintCmd,
		server.ExportCmd(ctx, cdc, appExport),
		client.LineBreak,
		version.VersionCmd,
	)
}
