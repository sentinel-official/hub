package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	hub "github.com/sentinel-official/sentinel-hub/types"
)

var _ sdk.Msg = (*MsgUpdateSessionInfo)(nil)

type MsgUpdateSessionInfo struct {
	From               sdk.AccAddress    `json:"from"`
	SubscriptionID     hub.ID            `json:"subscription_id"`
	Bandwidth          hub.Bandwidth     `json:"bandwidth"`
	NodeOwnerSignature auth.StdSignature `json:"node_owner_signature"`
	ClientSignature    auth.StdSignature `json:"client_signature"`
}

func (msg MsgUpdateSessionInfo) Type() string {
	return "MsgUpdateSessionInfo"
}

func (msg MsgUpdateSessionInfo) ValidateBasic() sdk.Error {
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

func (msg MsgUpdateSessionInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func (msg MsgUpdateSessionInfo) Route() string {
	return RouterKey
}

func NewMsgUpdateSessionInfo(from sdk.AccAddress,
	subscriptionID hub.ID, bandwidth hub.Bandwidth,
	nodeOwnerSignature, clientSignature auth.StdSignature) *MsgUpdateSessionInfo {

	return &MsgUpdateSessionInfo{
		From:               from,
		SubscriptionID:     subscriptionID,
		Bandwidth:          bandwidth,
		NodeOwnerSignature: nodeOwnerSignature,
		ClientSignature:    clientSignature,
	}
}
