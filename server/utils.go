package server

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	
	"github.com/sentinel-official/hub/version"
)

func AddCommands(ctx *server.Context, cdc *codec.Codec, root *cobra.Command,
	creator server.AppCreator, export server.AppExporter) {
	root.PersistentFlags().String("log_level", ctx.Config.LogLevel, "Log level")
	
	cmd := &cobra.Command{
		Use:   "tendermint",
		Short: "Tendermint subcommands",
	}
	
	cmd.AddCommand(
		server.ShowNodeIDCmd(ctx),
		server.ShowValidatorCmd(ctx),
		server.ShowAddressCmd(ctx),
		server.VersionCmd(ctx),
	)
	
	root.AddCommand(
		server.StartCmd(ctx, creator),
		server.UnsafeResetAllCmd(ctx),
		client.LineBreak,
		cmd,
		server.ExportCmd(ctx, cdc, export),
		client.LineBreak,
		version.Cmd,
	)
}
