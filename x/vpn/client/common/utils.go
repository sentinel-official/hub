package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func GetSubscriptionAndNodeByID(cliCtx context.CLIContext, cdc *codec.Codec,
	id sdkTypes.ID) (node vpn.Node, subscription vpn.Subscription, err error) {

	var res []byte

	res, err = QuerySubscription(cliCtx, cdc, id)
	if err != nil {
		return
	}
	if res == nil {
		err = fmt.Errorf("no subscription found")
		return
	}

	if err = cdc.UnmarshalJSON(res, &subscription); err != nil {
		return
	}

	res, err = QueryNode(cliCtx, cdc, subscription.NodeID)
	if err != nil {
		return
	}
	if res == nil {
		err = fmt.Errorf("no node found")
		return
	}

	if err = cdc.UnmarshalJSON(res, &node); err != nil {
		return
	}

	return node, subscription, nil
}

func BuildMsgUpdateSessionInfoAndSign(txBldr authTxBuilder.TxBuilder, cliCtx context.CLIContext, cdc *codec.Codec,
	id sdkTypes.ID, bandwidth sdkTypes.Bandwidth) (stdMsg authTxBuilder.StdSignMsg, sign auth.StdSignature, err error) {

	node, subscription, err := GetSubscriptionAndNodeByID(cliCtx, cdc, id)
	if err != nil {
		return stdMsg, sign, err
	}

	msg := vpn.NewMsgUpdateSessionInfo(node.Owner, subscription.Client, subscription.ID, bandwidth)

	txBldr, err = utils.PrepareTxBuilder(txBldr, cliCtx)
	if err != nil {
		return stdMsg, sign, err
	}

	stdSignMsg, err := txBldr.BuildSignMsg([]csdkTypes.Msg{msg})
	if err != nil {
		return stdMsg, sign, err
	}

	passphrase, err := keys.GetPassphrase(cliCtx.FromName)
	if err != nil {
		return stdMsg, sign, err
	}

	sign, err = authTxBuilder.MakeSignature(txBldr.Keybase(), cliCtx.FromName, passphrase, stdSignMsg)
	return stdSignMsg, sign, err
}
