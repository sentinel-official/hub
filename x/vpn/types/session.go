package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type SessionBandwidthInfo struct {
	ToProvide        sdkTypes.Bandwidth `json:"to_provide"`
	Consumed         sdkTypes.Bandwidth `json:"consumed"`
	NodeOwnerSign    []byte             `json:"node_owner_sign"`
	ClientSign       []byte             `json:"client_sign"`
	ModifiedAtHeight int64              `json:"modified_at_height"`
}

type Session struct {
	ID              sdkTypes.ID          `json:"id"`
	NodeID          sdkTypes.ID          `json:"node_id"`
	NodeOwner       csdkTypes.AccAddress `json:"node_owner"`
	NodeOwnerPubKey crypto.PubKey        `json:"node_owner_pub_key"`
	Client          csdkTypes.AccAddress `json:"client"`
	ClientPubKey    crypto.PubKey        `json:"client_pub_key"`
	Deposit         csdkTypes.Coin       `json:"deposit"`
	PricePerGB      csdkTypes.Coin       `json:"price_per_gb"`

	BandwidthInfo          SessionBandwidthInfo `json:"bandwidth_info"`
	Status                 string               `json:"status"`
	StatusModifiedAtHeight int64                `json:"status_modified_at_height"`
}

func (s Session) Amount() csdkTypes.Coin {
	consumedBandwidth := s.BandwidthInfo.Consumed.Upload.Add(s.BandwidthInfo.Consumed.Download)
	amountInt := consumedBandwidth.Mul(s.PricePerGB.Amount).Quo(sdkTypes.GB.Add(sdkTypes.GB))

	amount := csdkTypes.NewCoin(s.PricePerGB.Denom, amountInt)
	if s.Deposit.IsLT(amount) || s.Deposit.IsEqual(amount) {
		return s.Deposit
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
