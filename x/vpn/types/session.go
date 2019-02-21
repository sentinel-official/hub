package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type SessionBandwidth struct {
	ToProvide       sdkTypes.Bandwidth
	Consumed        sdkTypes.Bandwidth
	NodeOwnerSign   []byte
	ClientSign      []byte
	UpdatedAtHeight int64
}

type SessionDetails struct {
	ID              sdkTypes.ID
	NodeID          sdkTypes.ID
	NodeOwner       csdkTypes.AccAddress
	NodeOwnerPubKey crypto.PubKey
	Client          csdkTypes.AccAddress
	ClientPubKey    crypto.PubKey
	LockedAmount    csdkTypes.Coin
	PricePerGB      csdkTypes.Coin

	Bandwidth SessionBandwidth

	Status          string
	StatusAtHeight  int64
	StartedAtHeight int64
	EndedAtHeight   int64
}

func (s SessionDetails) Amount() csdkTypes.Coin {
	consumedBandwidth := s.Bandwidth.Consumed.Upload.Add(s.Bandwidth.Consumed.Download)
	amountInt := consumedBandwidth.Div(sdkTypes.GB.Add(sdkTypes.GB)).Mul(s.PricePerGB.Amount)

	amount := csdkTypes.NewCoin(s.PricePerGB.Denom, amountInt)
	if s.LockedAmount.IsLT(amount) || s.LockedAmount.IsEqual(amount) {
		return s.LockedAmount
	}

	return amount
}

func (s *SessionDetails) SetNewSessionBandwidth(signData *sdkTypes.BandwidthSignData,
	clientSign, nodeOwnerSign []byte, height int64) error {

	if signData.Bandwidth.LTE(s.Bandwidth.Consumed) ||
		s.Bandwidth.ToProvide.LT(signData.Bandwidth) {
		return errors.New("Invalid bandwidth")
	}

	signDataBytes := signData.GetBytes()
	if !s.ClientPubKey.VerifyBytes(signDataBytes, clientSign) ||
		!s.NodeOwnerPubKey.VerifyBytes(signDataBytes, nodeOwnerSign) {
		return errors.New("Invalid client sign or node owner sign")
	}

	s.Bandwidth.Consumed = signData.Bandwidth
	s.Bandwidth.ClientSign = clientSign
	s.Bandwidth.NodeOwnerSign = nodeOwnerSign
	s.Bandwidth.UpdatedAtHeight = height

	return nil
}
