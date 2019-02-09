package keeper

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func VerifyAndUpdateSessionBandwidth(ctx csdkTypes.Context, ak auth.AccountKeeper,
	session *types.SessionDetails, sign *types.BandwidthSign,
	clientSign, nodeOwnerSign []byte) csdkTypes.Error {

	if sign.Bandwidth.LTE(session.Bandwidth.Consumed) ||
		session.Bandwidth.ToProvide.LT(sign.Bandwidth) {
		return types.ErrorInvalidBandwidth()
	}

	signBytes, err := sign.GetBytes()
	if err != nil {
		return err
	}

	nodeOwner := ak.GetAccount(ctx, sign.NodeOwner)
	client := ak.GetAccount(ctx, sign.Client)

	if !client.GetPubKey().VerifyBytes(signBytes, clientSign) ||
		!nodeOwner.GetPubKey().VerifyBytes(signBytes, nodeOwnerSign) {
		return types.ErrorInvalidSign()
	}

	session.Bandwidth.Consumed = sign.Bandwidth
	session.Bandwidth.ClientSign = clientSign
	session.Bandwidth.NodeOwnerSign = nodeOwnerSign
	session.Bandwidth.UpdatedAtHeight = ctx.BlockHeight()
	if session.Status == types.StatusInit {
		session.StartedAtHeight = ctx.BlockHeight()
	}
	if session.Status != types.StatusActive {
		session.Status = types.StatusActive
		session.StatusAtHeight = ctx.BlockHeight()
	}

	return nil
}
