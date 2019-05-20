package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Subscription struct {
	ID                sdkTypes.ID          `json:"id"`
	NodeID            sdkTypes.ID          `json:"node_id"`
	Client            csdkTypes.AccAddress `json:"client"`
	ClientPubKey      crypto.PubKey        `json:"client_pub_key"`
	PricePerGB        csdkTypes.Coin       `json:"price_per_gb"`
	TotalDeposit      csdkTypes.Coin       `json:"total_deposit"`
	TotalBandwidth    sdkTypes.Bandwidth   `json:"total_bandwidth"`
	ConsumedDeposit   csdkTypes.Coin       `json:"consumed_deposit"`
	ConsumedBandwidth sdkTypes.Bandwidth   `json:"consumed_bandwidth"`
	SessionsCount     uint64               `json:"sessions_count"`
	Status            string               `json:"status"`
	StatusModifiedAt  int64                `json:"status_modified_at"`
}

func (s Subscription) String() string {
	clientPubKey, err := csdkTypes.Bech32ifyAccPub(s.ClientPubKey)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(`Subscription
  ID:                 %s
  NodeID:             %s
  Client Address:     %s
  Client Public Key:  %s
  Price Per GB:       %s
  Total Deposit:      %s
  Total Bandwidth:    %s
  Consumed Deposit:   %s
  Consumed Bandwidth: %s
  Sessions Count:     %d
  Status:             %s
  Status Modified At: %d`, s.ID, s.NodeID, s.Client, clientPubKey,
		s.PricePerGB, s.TotalDeposit, s.TotalBandwidth, s.ConsumedDeposit, s.ConsumedBandwidth,
		s.SessionsCount, s.Status, s.StatusModifiedAt)
}
