package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgRegisterNode)(nil)
)

type MsgRegisterNode struct {
	From           sdk.AccAddress     `json:"from"`
	Provider       hub.ProvAddress    `json:"provider"`
	PricePerGB     sdk.Coins          `json:"price_per_gb"`
	RemoteURL      string             `json:"remote_url"`
	Version        string             `json:"version"`
	BandwidthSpeed NodeBandwidthSpeed `json:"bandwidth_speed"`
	Category       NodeCategory       `json:"category"`
}

func NewMsgRegisterNode(from sdk.AccAddress, provider hub.ProvAddress, pricePerGB sdk.Coins,
	remoteURL, version string, bandwidthSpeed NodeBandwidthSpeed, category NodeCategory) MsgRegisterNode {
	return MsgRegisterNode{
		From:           from,
		Provider:       provider,
		PricePerGB:     pricePerGB,
		RemoteURL:      remoteURL,
		Version:        version,
		BandwidthSpeed: bandwidthSpeed,
		Category:       category,
	}
}

func (m MsgRegisterNode) Route() string {
	return RouterKey
}

func (m MsgRegisterNode) Type() string {
	return "register_node"
}

func (m MsgRegisterNode) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if m.Provider == nil || m.Provider.Empty() {
		return ErrorInvalidField("provider")
	}
	if m.PricePerGB == nil || !m.PricePerGB.IsValid() {
		return ErrorInvalidField("price_per_gb")
	}
	if len(m.RemoteURL) == 0 {
		return ErrorInvalidField("remove_url")
	}
	if len(m.Version) == 0 {
		return ErrorInvalidField("remote_url")
	}
	if m.BandwidthSpeed.IsAnyZero() {
		return ErrorInvalidField("bandwidth_speed")
	}
	if !m.Category.IsValid() {
		return ErrorInvalidField("category")
	}

	return nil
}

func (m MsgRegisterNode) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRegisterNode) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}
