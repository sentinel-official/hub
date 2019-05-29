package types

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

var _ csdkTypes.Msg = (*MsgUpdateSessionInfo)(nil)

type MsgUpdateSessionInfo struct {
	From               csdkTypes.AccAddress `json:"from"`
	SubscriptionID     sdkTypes.ID          `json:"subscription_id"`
	Bandwidth          sdkTypes.Bandwidth   `json:"bandwidth"`
	NodeOwnerSignature auth.StdSignature    `json:"node_owner_signature"`
	ClientSignature    auth.StdSignature    `json:"client_signature"`
}

func (msg MsgUpdateSessionInfo) Type() string {
	return "MsgUpdateSessionInfo"
}

func (msg MsgUpdateSessionInfo) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return ErrorInvalidField("from")
	}
	if !msg.Bandwidth.AllPositive() {
		return ErrorInvalidField("bandwidth")
	}
	if msg.NodeOwnerSignature.Signature == nil || msg.NodeOwnerSignature.PubKey == nil {
		return ErrorInvalidField("node_owner_signature")
	}
	if msg.ClientSignature.Signature == nil || msg.ClientSignature.PubKey == nil {
		return ErrorInvalidField("client_signature")
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
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgUpdateSessionInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionInfo(from csdkTypes.AccAddress,
	subscriptionID sdkTypes.ID, bandwidth sdkTypes.Bandwidth,
	nodeOwnerSignature, clientSignature auth.StdSignature) *MsgUpdateSessionInfo {

	return &MsgUpdateSessionInfo{
		From:               from,
		SubscriptionID:     subscriptionID,
		Bandwidth:          bandwidth,
		NodeOwnerSignature: nodeOwnerSignature,
		ClientSignature:    clientSignature,
	}
}
