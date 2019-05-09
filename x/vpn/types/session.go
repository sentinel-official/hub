package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type SessionBandwidthInfo struct {
	ToProvide        sdkTypes.Bandwidth
	Consumed         sdkTypes.Bandwidth
	NodeOwnerSign    []byte
	ClientSign       []byte
	ModifiedAtHeight int64
}

type Session struct {
	ID              sdkTypes.ID
	NodeID          sdkTypes.ID
	NodeOwner       csdkTypes.AccAddress
	NodeOwnerPubKey crypto.PubKey
	Client          csdkTypes.AccAddress
	ClientPubKey    crypto.PubKey
	DepositAmount   csdkTypes.Coin
	PricePerGB      csdkTypes.Coin

	BandwidthInfo          SessionBandwidthInfo
	StartedAtHeight        int64
	EndedAtHeight          int64
	Status                 string
	StatusModifiedAtHeight int64
}

func (s Session) Amount() csdkTypes.Coin {
	consumedBandwidth := s.BandwidthInfo.Consumed.Upload.Add(s.BandwidthInfo.Consumed.Download)
	amountInt := consumedBandwidth.Quo(sdkTypes.GB.Add(sdkTypes.GB)).Mul(s.PricePerGB.Amount)

	amount := csdkTypes.NewCoin(s.PricePerGB.Denom, amountInt)
	if s.DepositAmount.IsLT(amount) || s.DepositAmount.IsEqual(amount) {
		return s.DepositAmount
	}

	return amount
}

func (s *Session) UpdateSessionBandwidthInfo(consumed sdkTypes.Bandwidth,
	nodeOwnerSign, clientSign []byte, height int64) error {

	if consumed.LT(s.BandwidthInfo.Consumed) ||
		s.BandwidthInfo.ToProvide.LT(consumed) {
		return errors.New(errMsgInvalidBandwidth)
	}

	data := sdkTypes.NewBandwidthSign(s.ID, consumed, s.NodeOwner, s.Client).GetBytes()
	if !s.NodeOwnerPubKey.VerifyBytes(data, nodeOwnerSign) ||
		!s.ClientPubKey.VerifyBytes(data, clientSign) {
		return errors.New(errMsgInvalidBandwidthSigns)
	}

	s.BandwidthInfo.Consumed = consumed
	s.BandwidthInfo.NodeOwnerSign = nodeOwnerSign
	s.BandwidthInfo.ClientSign = clientSign
	s.BandwidthInfo.ModifiedAtHeight = height

	return nil
}
