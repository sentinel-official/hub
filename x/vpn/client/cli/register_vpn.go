package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	authCmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/ironman0x7b2/sentinel-hub/x/vpn"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagIP                = "ip"
	FlagPort              = "port"
	FlagUploadSpeed       = "upload-speed"
	FlagDownloadSpeed     = "download-speed"
	FlagPricePerGB        = "price-per-gb"
	FlagLocationLatitude  = "location-latitude"
	FlagLocationLongitude = "location-logitude"
	FlagLocationCity      = "location-city"
	FlagLocationCountry   = "location-country"
	FlagEncMethod         = "enc-method"
	FlagNodeType          = "node-type"
	FlagVersion           = "version"
)

func RegisterVpnCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register_vpn",
		Short: "Register for sentinel vpn service",
		RunE: func(cmd *cobra.Command, args []string) error {

			txBldr := authtxb.NewTxBuilderFromCLI().WithCodec(cdc)
			CliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(authCmd.GetAccountDecoder(cdc))
			ip := viper.GetString(FlagIP)
			port := viper.GetString(FlagPort)
			ppgb := viper.GetInt64(FlagPricePerGB)
			uploadSpeed := viper.GetInt64(FlagUploadSpeed)
			downloadSpeed := viper.GetInt64(FlagDownloadSpeed)
			latitude := viper.GetInt64(FlagLocationLatitude)
			longitude := viper.GetInt64(FlagLocationLongitude)
			city := viper.GetString(FlagLocationCity)
			country := viper.GetString(FlagLocationCountry)
			enc_method := viper.GetString(FlagEncMethod)
			node_type := viper.GetString(FlagNodeType)
			version := viper.GetString(FlagVersion)
			if err := CliCtx.EnsureAccountExists(); err != nil {
				return err
			}
			from, err := CliCtx.GetFromAddress()
			if err != nil {
				return err
			}
			account, err := CliCtx.GetAccount(from)
			if err != nil {
				return err
			}

			coins, err := sdkTypes.ParseCoins("100mycoin")
			if err != nil {
				return err
			}

			// ensure account has enough coins
			if !account.GetCoins().IsGTE(coins) {
				return errors.Errorf("Address %s doesn't have enough coins to pay for this transaction.", from)
			}

			msg := vpn.CreateRegisterVpnMsg(from, ip, port, coins, ppgb, uploadSpeed, downloadSpeed, latitude, longitude, city, country, enc_method, node_type, version)

			if CliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, CliCtx, []sdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, CliCtx, []sdkTypes.Msg{msg})
		},
	}
	cmd.Flags().String(FlagIP, "", "ip")
	cmd.Flags().String(FlagPort, "", "port")
	cmd.Flags().Int64(FlagUploadSpeed, -1, "upload_speed")
	cmd.Flags().Int64(FlagDownloadSpeed, -1, "download_speed")
	cmd.Flags().Int64(FlagPricePerGB, -1, "price_per_gb")
	cmd.Flags().String(FlagLocationLatitude, "", "location_latitude")
	cmd.Flags().String(FlagLocationLongitude, "", "location_longitude")
	cmd.Flags().String(FlagLocationCity, "", "location_city")
	cmd.Flags().String(FlagLocationCountry, "", "location_country")
	cmd.Flags().String(FlagEncMethod, "", "enc_method")
	cmd.Flags().String(FlagNodeType, "", "node type")
	cmd.Flags().String(FlagVersion, "", "version")
	return cmd
}
