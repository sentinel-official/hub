package types

import (
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
