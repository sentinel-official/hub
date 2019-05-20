package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var _ csdkTypes.Msg = (*MsgUpdateSessionInfo)(nil)

type MsgUpdateSessionInfo struct {
	From           csdkTypes.AccAddress `json:"from"`
	SubscriptionID sdkTypes.ID          `json:"subscription_id"`
	Bandwidth      sdkTypes.Bandwidth   `json:"bandwidth"`
	NodeSign       []byte               `json:"node_sign"`
	ClientSign     []byte               `json:"client_sign"`
}

func (msg MsgUpdateSessionInfo) Type() string {
	return "MsgUpdateSessionInfo"
}

func (msg MsgUpdateSessionInfo) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if msg.SubscriptionID == nil || msg.SubscriptionID.Len() == 0 {
		return ErrorInvalidField("subscription_id")
	}
	if !msg.Bandwidth.IsPositive() {
		return ErrorInvalidField("bandwidth")
	}
	if len(msg.NodeSign) == 0 {
		return ErrorInvalidField("node_sign")
	}
	if len(msg.ClientSign) == 0 {
		return ErrorInvalidField("client_sign")
	}

	return nil
}

func (msg MsgUpdateSessionInfo) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgUpdateSessionInfo) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateSessionInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionInfo(from csdkTypes.AccAddress,
	subscriptionID sdkTypes.ID, bandwidth sdkTypes.Bandwidth,
	nodeSign, clientSign []byte) *MsgUpdateSessionInfo {

	return &MsgUpdateSessionInfo{
		From:           from,
		SubscriptionID: subscriptionID,
		Bandwidth:      bandwidth,
		NodeSign:       nodeSign,
		ClientSign:     clientSign,
	}
}
