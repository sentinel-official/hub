package cli

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptoKeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authTxBuiler "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmConfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
	tmCli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/sentinel-official/sentinel-hub/app"
)

// nolint:gochecknoglobals
var (
	defaultTokens                  = sdk.TokensFromTendermintPower(100)
	defaultAmount                  = defaultTokens.String() + "stake"
	defaultCommissionRate          = "0.1"
	defaultCommissionMaxRate       = "0.2"
	defaultCommissionMaxChangeRate = "0.01"
	defaultMinSelfDelegation       = "1"
)

// nolint:gocyclo
func GenTxCmd(ctx *server.Context, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gentx",
		Short: "Generate a genesis tx carrying a self delegation",
		Args:  cobra.NoArgs,
		Long: fmt.Sprintf(`This command is an alias of the 'sentinel-hubd tx create-validator' command'.

It creates a genesis piece carrying a self delegation with the
following delegation and commission default parameters:

	delegation amount:           %s
	commission rate:             %s
	commission max rate:         %s
	commission max change rate:  %s
	minimum self delegation:     %s
`, defaultAmount, defaultCommissionRate, defaultCommissionMaxRate,
			defaultCommissionMaxChangeRate, defaultMinSelfDelegation),
		RunE: func(cmd *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(tmCli.HomeFlag))

			nodeID, valPubKey, err := InitializeNodeValidatorFiles(ctx.Config)
			if err != nil {
				return err
			}

			if nodeIDString := viper.GetString(cli.FlagNodeID); nodeIDString != "" {
				nodeID = nodeIDString
			}

			ip := viper.GetString(cli.FlagIP)
			if ip == "" {
				_, _ = fmt.Fprintf(os.Stderr, "couldn't retrieve an external IP; "+
					"the tx's memo field will be unset")
			}

			genDoc, err := loadGenesisDoc(cdc, config.GenesisFile())
			if err != nil {
				return err
			}

			state := app.GenesisState{}
			if err = cdc.UnmarshalJSON(genDoc.AppState, &state); err != nil {
				return err
			}

			if err = app.ValidateGenesisState(state); err != nil {
				return err
			}

			kb, err := keys.NewKeyBaseFromDir(viper.GetString(flagClientHome))
			if err != nil {
				return err
			}

			name := viper.GetString(client.FlagName)
			key, err := kb.Get(name)
			if err != nil {
				return err
			}

			if valPubKeyString := viper.GetString(cli.FlagPubKey); valPubKeyString != "" {
				valPubKey, err = sdk.GetConsPubKeyBech32(valPubKeyString)
				if err != nil {
					return err
				}
			}

			website := viper.GetString(cli.FlagWebsite)
			details := viper.GetString(cli.FlagDetails)
			identity := viper.GetString(cli.FlagIdentity)

			prepareFlagsForTxCreateValidator(config, nodeID, ip, genDoc.ChainID, valPubKey, website, details, identity)

			amount := viper.GetString(cli.FlagAmount)
			coins, err := sdk.ParseCoins(amount)
			if err != nil {
				return err
			}

			err = accountInGenesis(state, key.GetAddress(), coins)
			if err != nil {
				return err
			}

			txBuilder := authTxBuiler.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			viper.Set(client.FlagGenerateOnly, true)

			txBuilder, msg, err := cli.BuildCreateValidatorMsg(cliCtx, txBuilder)
			if err != nil {
				return err
			}

			info, err := txBuilder.Keybase().Get(name)
			if err != nil {
				return err
			}

			if info.GetType() == cryptoKeys.TypeOffline || info.GetType() == cryptoKeys.TypeMulti {
				fmt.Println("Offline key passed in. Use `sentinel-hubcli tx sign` command to sign:")
				return utils.PrintUnsignedStdTx(txBuilder, cliCtx, []sdk.Msg{msg}, true)
			}

			w := bytes.NewBuffer([]byte{})
			cliCtx = cliCtx.WithOutput(w)
			if err = utils.PrintUnsignedStdTx(txBuilder, cliCtx, []sdk.Msg{msg}, true); err != nil {
				return err
			}

			stdTx, err := readUnsignedGenTxFile(cdc, w)
			if err != nil {
				return err
			}

			signedTx, err := utils.SignStdTx(txBuilder, cliCtx, name, stdTx, false, true)
			if err != nil {
				return err
			}

			outputDocument := viper.GetString(client.FlagOutputDocument)
			if outputDocument == "" {
				outputDocument, err = makeOutputFilepath(config.RootDir, nodeID)
				if err != nil {
					return err
				}
			}

			if err := writeSignedGenTx(cdc, outputDocument, signedTx); err != nil {
				return err
			}

			_, _ = fmt.Fprintf(os.Stderr, "Genesis transaction written to %q\n", outputDocument)
			return nil
		},
	}

	ip, _ := server.ExternalIP()

	cmd.Flags().String(tmCli.HomeFlag, app.DefaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, app.DefaultCLIHome, "client's home directory")
	cmd.Flags().String(client.FlagName, "", "name of private key with which to sign the gentx")
	cmd.Flags().String(client.FlagOutputDocument, "",
		"write the genesis transaction JSON document to the given file instead of the default location")
	cmd.Flags().String(cli.FlagIP, ip, "The node's public IP")
	cmd.Flags().String(cli.FlagNodeID, "", "The node's NodeID")
	cmd.Flags().String(cli.FlagWebsite, "", "The validator's (optional) website")
	cmd.Flags().String(cli.FlagDetails, "", "The validator's (optional) details")
	cmd.Flags().String(cli.FlagIdentity, "", "The (optional) identity signature (ex. UPort or Keybase)")
	cmd.Flags().AddFlagSet(cli.FsCommissionCreate)
	cmd.Flags().AddFlagSet(cli.FsMinSelfDelegation)
	cmd.Flags().AddFlagSet(cli.FsAmount)
	cmd.Flags().AddFlagSet(cli.FsPk)

	_ = cmd.MarkFlagRequired(client.FlagName)

	return cmd
}

func accountInGenesis(state app.GenesisState, key sdk.AccAddress, coins sdk.Coins) error {
	accountIsInGenesis := false
	bondDenom := state.Staking.Params.BondDenom

	for _, acc := range state.Accounts {
		if acc.Address.Equals(key) {
			if coins.AmountOf(bondDenom).GT(acc.Coins.AmountOf(bondDenom)) {
				return fmt.Errorf(
					"account %v is in genesis, but it only has %v%v available to stake, not %v%v",
					key.String(), acc.Coins.AmountOf(bondDenom), bondDenom, coins.AmountOf(bondDenom), bondDenom,
				)
			}

			accountIsInGenesis = true
			break
		}
	}

	if accountIsInGenesis {
		return nil
	}

	return fmt.Errorf("account %s in not in the app_state.accounts array of genesis.json", key)
}

func prepareFlagsForTxCreateValidator(config *tmConfig.Config, nodeID, ip, chainID string,
	valPubKey crypto.PubKey, website, details, identity string) {
	viper.Set(tmCli.HomeFlag, viper.GetString(flagClientHome))
	viper.Set(client.FlagChainID, chainID)
	viper.Set(client.FlagFrom, viper.GetString(client.FlagName))
	viper.Set(cli.FlagNodeID, nodeID)
	viper.Set(cli.FlagIP, ip)
	viper.Set(cli.FlagPubKey, sdk.MustBech32ifyConsPub(valPubKey))
	viper.Set(cli.FlagMoniker, config.Moniker)
	viper.Set(cli.FlagWebsite, website)
	viper.Set(cli.FlagDetails, details)
	viper.Set(cli.FlagIdentity, identity)

	if config.Moniker == "" {
		viper.Set(cli.FlagMoniker, viper.GetString(client.FlagName))
	}
	if viper.GetString(cli.FlagAmount) == "" {
		viper.Set(cli.FlagAmount, defaultAmount)
	}
	if viper.GetString(cli.FlagCommissionRate) == "" {
		viper.Set(cli.FlagCommissionRate, defaultCommissionRate)
	}
	if viper.GetString(cli.FlagCommissionMaxRate) == "" {
		viper.Set(cli.FlagCommissionMaxRate, defaultCommissionMaxRate)
	}
	if viper.GetString(cli.FlagCommissionMaxChangeRate) == "" {
		viper.Set(cli.FlagCommissionMaxChangeRate, defaultCommissionMaxChangeRate)
	}
	if viper.GetString(cli.FlagMinSelfDelegation) == "" {
		viper.Set(cli.FlagMinSelfDelegation, defaultMinSelfDelegation)
	}
}

func makeOutputFilepath(rootDir, nodeID string) (string, error) {
	writePath := filepath.Join(rootDir, "config", "gentx")
	if err := common.EnsureDir(writePath, 0700); err != nil {
		return "", err
	}

	return filepath.Join(writePath, fmt.Sprintf("gentx-%v.json", nodeID)), nil
}

func readUnsignedGenTxFile(cdc *codec.Codec, r io.Reader) (auth.StdTx, error) {
	var stdTx auth.StdTx

	bz, err := ioutil.ReadAll(r)
	if err != nil {
		return stdTx, err
	}

	err = cdc.UnmarshalJSON(bz, &stdTx)
	return stdTx, err
}

func writeSignedGenTx(cdc *codec.Codec, outputDocument string, tx auth.StdTx) error {
	outputFile, err := os.OpenFile(outputDocument, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func() {
		if err = outputFile.Close(); err != nil {
			panic(err)
		}
	}()

	json, err := cdc.MarshalJSON(tx)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(outputFile, "%s\n", json)
	return err
}
