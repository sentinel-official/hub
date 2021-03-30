package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgSubscribeToNodeRequest)(nil)
	_ sdk.Msg = (*MsgSubscribeToPlanRequest)(nil)
	_ sdk.Msg = (*MsgCancelRequest)(nil)

	_ sdk.Msg = (*MsgAddQuotaRequest)(nil)
	_ sdk.Msg = (*MsgUpdateQuotaRequest)(nil)
)

func NewMsgSubscribeToNodeRequest(from, address string, deposit sdk.Coin) MsgSubscribeToNodeRequest {
	return MsgSubscribeToNodeRequest{
		From:    from,
		Address: address,
		Deposit: deposit,
	}
}

func (m MsgSubscribeToNodeRequest) Route() string {
	return RouterKey
}

func (m MsgSubscribeToNodeRequest) Type() string {
	return fmt.Sprintf("%s:subscribe_to_node", ModuleName)
}

func (m MsgSubscribeToNodeRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Address should be valid
	if _, err := hub.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	// Deposit should be valid and positive
	if !m.Deposit.IsValid() || !m.Deposit.IsPositive() {
		return errors.Wrapf(ErrorInvalidField, "%s", "deposit")
	}

	return nil
}

func (m MsgSubscribeToNodeRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSubscribeToNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgSubscribeToPlanRequest(from string, id uint64, denom string) MsgSubscribeToPlanRequest {
	return MsgSubscribeToPlanRequest{
		From:  from,
		Id:    id,
		Denom: denom,
	}
}

func (m MsgSubscribeToPlanRequest) Route() string {
	return RouterKey
}

func (m MsgSubscribeToPlanRequest) Type() string {
	return fmt.Sprintf("%s:subscribe_to_plan", ModuleName)
}

func (m MsgSubscribeToPlanRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Denom should be valid
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "denom")
	}

	return nil
}

func (m MsgSubscribeToPlanRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSubscribeToPlanRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgCancelRequest(from string, id uint64) MsgCancelRequest {
	return MsgCancelRequest{
		From: from,
		Id:   id,
	}
}

func (m MsgCancelRequest) Route() string {
	return RouterKey
}

func (m MsgCancelRequest) Type() string {
	return fmt.Sprintf("%s:cancel", ModuleName)
}

func (m MsgCancelRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	return nil
}

func (m MsgCancelRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgCancelRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgAddQuotaRequest(from string, id uint64, address string, bytes sdk.Int) MsgAddQuotaRequest {
	return MsgAddQuotaRequest{
		From:    from,
		Id:      id,
		Address: address,
		Bytes:   bytes,
	}
}

func (m MsgAddQuotaRequest) Route() string {
	return RouterKey
}

func (m MsgAddQuotaRequest) Type() string {
	return fmt.Sprintf("%s:add_quota", ModuleName)
}

func (m MsgAddQuotaRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address should be valid
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	// Bytes should be positive
	if !m.Bytes.IsPositive() {
		return errors.Wrapf(ErrorInvalidField, "%s", "bytes")
	}

	return nil
}

func (m MsgAddQuotaRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddQuotaRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateQuotaRequest(from string, id uint64, address string, bytes sdk.Int) MsgUpdateQuotaRequest {
	return MsgUpdateQuotaRequest{
		From:    from,
		Id:      id,
		Address: address,
		Bytes:   bytes,
	}
}

func (m MsgUpdateQuotaRequest) Route() string {
	return RouterKey
}

func (m MsgUpdateQuotaRequest) Type() string {
	return fmt.Sprintf("%s:update_quota", ModuleName)
}

func (m MsgUpdateQuotaRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address shouldn be valid
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	// Bytes should be positive
	if !m.Bytes.IsPositive() {
		return errors.Wrapf(ErrorInvalidField, "%s", "bytes")
	}

	return nil
}

func (m MsgUpdateQuotaRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdateQuotaRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
