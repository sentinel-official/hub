package cli

import (
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	ckeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func RegisterCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register Sentinel VPN service node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			apiPort := viper.GetInt64(flagAPIPort)
			latitude := viper.GetInt64(flagLocationLatitude)
			longitude := viper.GetInt64(flagLocationLongitude)
			city := viper.GetString(flagLocationCity)
			country := viper.GetString(flagLocationCountry)
			upload := viper.GetInt64(flagUploadSpeed)
			download := viper.GetInt64(flagDownloadSpeed)
			encMethod := viper.GetString(flagEncMethod)
			pricePerGB := viper.GetInt64(flagPricePerGB)
			version := viper.GetString(flagVersion)
			amount := viper.GetString(flagAmount)

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

			coins, err := csdkTypes.ParseCoins(amount)

			if err != nil {
				return err
			}

			if !account.GetCoins().IsAllGTE(coins) {
				return errors.Errorf("Address %s doesn't have enough coins to pay for this transaction.", from)
			}

			vpnsCountBytes, err := cliCtx.QueryStore(cdc.MustMarshalBinaryLengthPrefixed(sdkTypes.VPNsCountKey(from)), sdkTypes.KeyVPN)

			if err != nil {
				return err
			}

			var vpnsCount uint64

			if vpnsCountBytes != nil {
				if err := cdc.UnmarshalBinaryLengthPrefixed(vpnsCountBytes, &vpnsCount); err != nil {
					return err
				}
			}

			kb, err := ckeys.GetKeyBase()

			if err != nil {
				return err
			}

			name, err := cliCtx.GetFromName()

			if err != nil {
				return err
			}

			keyInfo, err := kb.Get(name)

			if err != nil {
				return err
			}

			pubKey := keyInfo.GetPubKey()
			lockerID := sdkTypes.KeyVPN + "/" + from.String() + "/" + strconv.Itoa(int(vpnsCount))
			msgLockerCoins := hub.MsgLockCoins{
				LockerID: lockerID,
				Coins:    coins,
				PubKey:   pubKey,
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
				return utils.PrintUnsignedStdTx(os.Stdout, txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().Int64(flagAPIPort, 3000, "Node API port")
	cmd.Flags().Int64(flagLocationLatitude, 0, "Latitude")
	cmd.Flags().Int64(flagLocationLongitude, 0, "Longitude")
	cmd.Flags().String(flagLocationCity, "", "City name")
	cmd.Flags().String(flagLocationCountry, "", "Country name")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet download speed in bytes/sec")
	cmd.Flags().String(flagEncMethod, "", "VPN tunnel encryption method")
	cmd.Flags().Int64(flagPricePerGB, 0, "Usage price per 1 GB in sent")
	cmd.Flags().String(flagVersion, "", "Node version")
	cmd.Flags().String(flagAmount, "100sent", "Amount of coins that you want to lock")

	return cmd
}
