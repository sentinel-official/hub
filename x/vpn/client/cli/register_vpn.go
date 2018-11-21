package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	ckeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	flagApiPort           = "api-port"
	flagVPNPort           = "vpn-port"
	flagAmount            = "amount"
	flagUploadSpeed       = "upload"
	flagDownloadSpeed     = "download"
	flagPricePerGB        = "price-per-gb"
	flagLocationLatitude  = "latitude"
	flagLocationLongitude = "longitude"
	flagLocationCity      = "city"
	flagLocationCountry   = "country"
	flagEncMethod         = "enc-method"
	flagVersion           = "version"
)

func RegisterVPNCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-vpn",
		Short: "Register for sentinel vpn service",
		RunE: func(cmd *cobra.Command, args []string) error {

			var kb keys.Keybase
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(authCli.GetAccountDecoder(cdc))

			apiPort := viper.GetString(flagApiPort)
			vpnPort := viper.GetString(flagVPNPort)
			amount := viper.GetString(flagAmount)
			pricePerGb := viper.GetInt64(flagPricePerGB)
			upload := viper.GetInt64(flagUploadSpeed)
			download := viper.GetInt64(flagDownloadSpeed)
			latitude := viper.GetInt64(flagLocationLatitude)
			longitude := viper.GetInt64(flagLocationLongitude)
			city := viper.GetString(flagLocationCity)
			country := viper.GetString(flagLocationCountry)
			encMethod := viper.GetString(flagEncMethod)
			version := viper.GetString(flagVersion)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			from, err := cliCtx.GetFromAddress()

			if err != nil {
				return err
			}

			account, err := cliCtx.GetAccount(from)

			if err != nil {
				return err
			}

			pubkey := account.GetPubKey()

			coins, err := csdkTypes.ParseCoins(amount)

			if err != nil {
				return err
			}

			// ensure account has enough coins
			if !account.GetCoins().IsGTE(coins) {
				return errors.Errorf("Address %s doesn't have enough coins to pay for this transaction.", from)
			}

			sequence, err := cliCtx.GetAccountSequence(from)

			if err != nil {
				return err
			}

			unSignBytes := sdkTypes.GetUnSignBytes(from, sequence, coins, pubkey)

			kb, err = ckeys.GetKeyBase()

			name, err := cliCtx.GetFromName()

			if err != nil {
				return err
			}

			passPhrase, err := ckeys.GetPassphrase(name)

			if err != nil {
				return err
			}

			signature, _, err := kb.Sign(name, passPhrase, unSignBytes)

			if err != nil {
				return err
			}

			msg := vpn.NewRegisterVPNMsg(from, coins,
				apiPort, vpnPort, pubkey,
				upload, download,
				latitude, longitude, city, country,
				pricePerGb, encMethod, version, sequence, signature)

			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagApiPort, "", "api port")
	cmd.Flags().String(flagVPNPort, "", " vpn port")
	cmd.Flags().String(flagAmount, "100mycoin", "amount")
	cmd.Flags().Int64(flagUploadSpeed, -1, "upload_speed")
	cmd.Flags().Int64(flagDownloadSpeed, -1, "download_speed")
	cmd.Flags().Int64(flagPricePerGB, -1, "price_per_gb")
	cmd.Flags().String(flagLocationLatitude, "", "location_latitude")
	cmd.Flags().String(flagLocationLongitude, "", "location_longitude")
	cmd.Flags().String(flagLocationCity, "", "location_city")
	cmd.Flags().String(flagLocationCountry, "", "location_country")
	cmd.Flags().String(flagEncMethod, "", "enc_method")
	cmd.Flags().String(flagVersion, "", "version")

	return cmd
}
