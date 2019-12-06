package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgRegisterResolver struct {
	From       sdk.AccAddress
	Commission sdk.Dec
}

func (msg MsgRegisterResolver) Route() string {
	return RouterKey
}

func (msg MsgRegisterResolver) Type() string {
	return "register_resolver_node"
}

func (msg MsgRegisterResolver) ValidateBasic() sdk.Error {
	return nil //TODO
}

func (msg MsgRegisterResolver) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return bz
}

func (msg MsgRegisterResolver) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}
func NewMsgRegisterResolver(from sdk.AccAddress, commission sdk.Dec) MsgRegisterResolver {
	return MsgRegisterResolver{
		From:       from,
		Commission: commission,
	}
}

var _ sdk.Msg = (*MsgRegisterResolver)(nil)

type MsgUpdateResolverInfo struct {
	From       sdk.AccAddress
	Commission sdk.Dec
}

func (msg MsgUpdateResolverInfo) Route() string {
	return RouterKey
}

func (msg MsgUpdateResolverInfo) Type() string {
	return "update_resolver_info"
}

func (msg MsgUpdateResolverInfo) ValidateBasic() sdk.Error {
	return nil //TODO
}

func (msg MsgUpdateResolverInfo) GetSignBytes() []byte {
	bz, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return bz
}

func (msg MsgUpdateResolverInfo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}

func NewMsgUpdateResolverInfo(from sdk.AccAddress, commission sdk.Dec) MsgUpdateResolverInfo {
	return MsgUpdateResolverInfo{
		From:       from,
		Commission: commission,
	}
}

var _ sdk.Msg = (*MsgUpdateResolverInfo)(nil)
