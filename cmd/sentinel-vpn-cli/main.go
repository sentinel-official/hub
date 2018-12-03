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

	app "github.com/ironman0x7b2/sentinel-sdk/apps/sentinel-vpn"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	ibcCli "github.com/ironman0x7b2/sentinel-sdk/x/ibc/client/cli"
	vpnCli "github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/cli"
)

var rootCmd = &cobra.Command{
	Use:   "sentinel-vpn-cli",
	Short: "Sentinel VPN light-client",
}

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	config := csdkTypes.GetConfig()
	config.SetBech32PrefixForAccount("cosmos", "cosmospub")
	config.SetBech32PrefixForValidator("cosmosvaloper", "cosmosvaloperpub")
	config.SetBech32PrefixForConsensusNode("cosmosvalcons", "cosmosvalconspub")
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
		stakeCli.GetCmdQueryValidator(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryValidators(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryValidatorUnbondingDelegations(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryValidatorRedelegations(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryDelegation(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryDelegations(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryPool(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryParams(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryUnbondingDelegation(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryUnbondingDelegations(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryRedelegation(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdQueryRedelegations(sdkTypes.KeyStake, cdc),
		slashingCli.GetCmdQuerySigningInfo(sdkTypes.KeySlashing, cdc),
		stakeCli.GetCmdQueryValidatorDelegations(sdkTypes.KeyStake, cdc),
		authCli.GetAccountCmd(sdkTypes.KeyAccount, cdc),
	)

	rootCmd.AddCommand(
		bankCli.SendTxCmd(cdc),
		cIBCCli.IBCTransferCmd(cdc),
		stakeCli.GetCmdCreateValidator(cdc),
		stakeCli.GetCmdEditValidator(cdc),
		stakeCli.GetCmdDelegate(cdc),
		stakeCli.GetCmdUnbond(sdkTypes.KeyStake, cdc),
		stakeCli.GetCmdRedelegate(sdkTypes.KeyStake, cdc),
		slashingCli.GetCmdUnjail(cdc),
	)

	rootCmd.AddCommand(client.LineBreak)
	rootCmd.AddCommand(
		client.PostCommands(
			ibcCli.IBCRelayCmd(cdc, sdkTypes.KeyIBC, sdkTypes.KeyAccount),
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
	authRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, sdkTypes.KeyAccount)
	bankRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	stakeRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	slashingRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
}
