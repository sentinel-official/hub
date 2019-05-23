package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var _ csdkTypes.Msg = (*MsgUpdateSessionInfo)(nil)

type MsgUpdateSessionInfo struct {
	From           csdkTypes.AccAddress `json:"from"`
	Client         csdkTypes.AccAddress `json:"client"`
	SubscriptionID sdkTypes.ID          `json:"subscription_id"`
	Bandwidth      sdkTypes.Bandwidth   `json:"bandwidth"`
}

func (msg MsgUpdateSessionInfo) Type() string {
	return "MsgUpdateSessionInfo"
}

func (msg MsgUpdateSessionInfo) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.Client == nil || msg.Client.Empty() {
		return ErrorInvalidField("client")
	}
	if !msg.Bandwidth.AllPositive() {
		return ErrorInvalidField("bandwidth")
	}

	return nil
}

func (msg MsgUpdateSessionInfo) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgUpdateSessionInfo) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From, msg.Client}
}

func (msg MsgUpdateSessionInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionInfo(from, client csdkTypes.AccAddress,
	subscriptionID sdkTypes.ID, bandwidth sdkTypes.Bandwidth) *MsgUpdateSessionInfo {

	return &MsgUpdateSessionInfo{
		From:           from,
		Client:         client,
		SubscriptionID: subscriptionID,
		Bandwidth:      bandwidth,
	}
}
