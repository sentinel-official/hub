package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	_ "github.com/cosmos/cosmos-sdk/client/lcd/statik"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authRest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	bankRest "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	cIBCCli "github.com/cosmos/cosmos-sdk/x/ibc/client/cli"
	slashingCli "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
	slashingRest "github.com/cosmos/cosmos-sdk/x/slashing/client/rest"
	stakeCli "github.com/cosmos/cosmos-sdk/x/stake/client/cli"
	stakeRest "github.com/cosmos/cosmos-sdk/x/stake/client/rest"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"

	ibcCli "github.com/ironman0x7b2/sentinel-sdk/x/ibc/client/cli"

	app "github.com/ironman0x7b2/sentinel-sdk/apps/sentinel-vpn"
	vpnCli "github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/cli"
)

const (
	storeAcc      = "acc"
	storeSlashing = "slashing"
	storeStake    = "stake"
)

var rootCmd = &cobra.Command{
	Use:   "sentinel-vpn-cli",
	Short: "Sentinel VPN light-client",
}

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	config := csdkTypes.GetConfig()
	config.SetBech32PrefixForAccount("sentacc", "sentpub")
	config.SetBech32PrefixForValidator("sentval", "sentvalpub")
	config.SetBech32PrefixForConsensusNode("sentcons", "sentconspub")
	config.Seal()

	rootCmd.AddCommand(
		rpc.InitClientCommand(),
		rpc.StatusCommand(),
		client.LineBreak,
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
		client.LineBreak,
	)

	rootCmd.AddCommand(
		stakeCli.GetCmdQueryValidator(storeStake, cdc),
		stakeCli.GetCmdQueryValidators(storeStake, cdc),
		stakeCli.GetCmdQueryValidatorUnbondingDelegations(storeStake, cdc),
		stakeCli.GetCmdQueryValidatorRedelegations(storeStake, cdc),
		stakeCli.GetCmdQueryDelegation(storeStake, cdc),
		stakeCli.GetCmdQueryDelegations(storeStake, cdc),
		stakeCli.GetCmdQueryPool(storeStake, cdc),
		stakeCli.GetCmdQueryParams(storeStake, cdc),
		stakeCli.GetCmdQueryUnbondingDelegation(storeStake, cdc),
		stakeCli.GetCmdQueryUnbondingDelegations(storeStake, cdc),
		stakeCli.GetCmdQueryRedelegation(storeStake, cdc),
		stakeCli.GetCmdQueryRedelegations(storeStake, cdc),
		slashingCli.GetCmdQuerySigningInfo(storeSlashing, cdc),
		stakeCli.GetCmdQueryValidatorDelegations(storeStake, cdc),
		authCli.GetAccountCmd(storeAcc, cdc),
	)

	rootCmd.AddCommand(
		bankCli.SendTxCmd(cdc),
		cIBCCli.IBCTransferCmd(cdc),
		ibcCli.IBCRelayCmd(cdc),
		stakeCli.GetCmdCreateValidator(cdc),
		stakeCli.GetCmdEditValidator(cdc),
		stakeCli.GetCmdDelegate(cdc),
		stakeCli.GetCmdUnbond(storeStake, cdc),
		stakeCli.GetCmdRedelegate(storeStake, cdc),
		slashingCli.GetCmdUnjail(cdc),
	)

	rootCmd.AddCommand(client.LineBreak)
	rootCmd.AddCommand(
		client.PostCommands(
			vpnCli.RegisterCommand(cdc),
			vpnCli.PaymentCommand(cdc),
			vpnCli.UpdateSessionStatusCommand(cdc),
			vpnCli.DeregisterCommand(cdc),
		)...)

	rootCmd.AddCommand(
		client.LineBreak,
		lcd.ServeCommand(cdc, registerRoutes),
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	executor := cli.PrepareMainCmd(rootCmd, "SV", app.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func registerRoutes(rs *lcd.RestServer) {
	keys.RegisterRoutes(rs.Mux, rs.CliCtx.Indent)
	rpc.RegisterRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	authRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, storeAcc)
	bankRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	stakeRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	slashingRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
}
