package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	ckeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

const (
	flagAPIPort           = "api-port"
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

			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(authCli.GetAccountDecoder(cdc))

			apiPort := viper.GetString(flagAPIPort)
			amount := viper.GetString(flagAmount)
			pricePerGB := viper.GetInt64(flagPricePerGB)
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

			pubKey := account.GetPubKey()
			coins, err := csdkTypes.ParseCoins(amount)

			if err != nil {
				return err
			}

			if !account.GetCoins().IsAllGTE(coins) {
				return errors.Errorf("Address %s doesn't have enough coins to pay for this transaction.", from)
			}

			sequence, err := cliCtx.GetAccountSequence(from)

			if err != nil {
				return err
			}

			lockerID := "vpn" + "/" + from.String() + "/" + strconv.Itoa(int(sequence)+1)
			msgLockerCoins := hub.MsgLockCoins{
				LockerID: lockerID,
				Coins:    coins,
				PubKey:   pubKey,
			}
			kb, err := ckeys.GetKeyBase()

			if err != nil {
				return err
			}

			name, err := cliCtx.GetFromName()

			if err != nil {
				return err
			}

			passPhrase, err := ckeys.GetPassphrase(name)

			if err != nil {
				return err
			}

			signature, _, err := kb.Sign(name, passPhrase, msgLockerCoins.GetUnSignBytes())

			if err != nil {
				return err
			}

			msg := vpn.NewMsgRegisterNode(from, apiPort,
				latitude, longitude, city, country,
				upload, download,
				encMethod, pricePerGB, version,
				lockerID, coins, pubKey, signature)

			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagAPIPort, "", "api port")
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
