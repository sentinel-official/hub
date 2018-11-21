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

func PayVPNServiceCommand(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay-vpn",
		Short: "pay for vpn",
		RunE: func(cmd *cobra.Command, args []string) error {

			var kb keys.Keybase
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithCodec(cdc)
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(authCli.GetAccountDecoder(cdc))

			vpnID := viper.GetString(flagVPNID)
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

			// ensure account has enough coins
			if !account.GetCoins().IsGTE(coins) {
				return errors.Errorf("Address %s doesn't have enough coins to pay for this transaction.", from)
			}

			sequence, err := cliCtx.GetAccountSequence(from)

			if err != nil {
				return err
			}

			pubkey := account.GetPubKey()

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

			msg := vpn.NewMsgPayVPNService(coins, vpnID, from, sequence, pubkey, signature)
			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagVPNID, "", "VPN id")
	cmd.Flags().String(flagAmount, "", "Amount of coins")

	return cmd

}
