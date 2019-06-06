package types

import (
	"encoding/json"

	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

var _ csdk.Msg = (*MsgUpdateSessionInfo)(nil)

type MsgUpdateSessionInfo struct {
	From               csdk.AccAddress   `json:"from"`
	SubscriptionID     sdk.ID            `json:"subscription_id"`
	Bandwidth          sdk.Bandwidth     `json:"bandwidth"`
	NodeOwnerSignature auth.StdSignature `json:"node_owner_signature"`
	ClientSignature    auth.StdSignature `json:"client_signature"`
}

func (msg MsgUpdateSessionInfo) Type() string {
	return "MsgUpdateSessionInfo"
}

func (msg MsgUpdateSessionInfo) ValidateBasic() csdk.Error {
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

func (msg MsgUpdateSessionInfo) GetSigners() []csdk.AccAddress {
	return []csdk.AccAddress{msg.From}
}

func (msg MsgUpdateSessionInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionInfo(from csdk.AccAddress,
	subscriptionID sdk.ID, bandwidth sdk.Bandwidth,
	nodeOwnerSignature, clientSignature auth.StdSignature) *MsgUpdateSessionInfo {

	return &MsgUpdateSessionInfo{
		From:               from,
		SubscriptionID:     subscriptionID,
		Bandwidth:          bandwidth,
		NodeOwnerSignature: nodeOwnerSignature,
		ClientSignature:    clientSignature,
	}
}
