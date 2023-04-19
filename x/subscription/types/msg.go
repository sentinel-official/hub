package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgCancelRequest)(nil)
	_ sdk.Msg = (*MsgShareRequest)(nil)
	_ sdk.Msg = (*MsgUpdateQuotaRequest)(nil)
)

func NewMsgCancelRequest(from sdk.AccAddress, id uint64) *MsgCancelRequest {
	return &MsgCancelRequest{
		From: from.String(),
		Id:   id,
	}
}

func (m *MsgCancelRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}

	return nil
}

func (m *MsgCancelRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgShareRequest(from sdk.AccAddress, id uint64, address sdk.AccAddress, bytes sdk.Int) *MsgShareRequest {
	return &MsgShareRequest{
		From:    from.String(),
		Id:      id,
		Address: address.String(),
		Bytes:   bytes,
	}
}

func (m *MsgShareRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidAddress, "address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidAddress, "%s", err)
	}
	if m.Bytes.IsNegative() {
		return errors.Wrap(ErrorInvalidBytes, "bytes cannot be negative")
	}

	return nil
}

func (m *MsgShareRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateQuotaRequest(from sdk.AccAddress, id uint64, address sdk.AccAddress, bytes sdk.Int) *MsgUpdateQuotaRequest {
	return &MsgUpdateQuotaRequest{
		From:    from.String(),
		Id:      id,
		Address: address.String(),
		Bytes:   bytes,
	}
}

func (m *MsgUpdateQuotaRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidAddress, "address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidAddress, "%s", err)
	}
	if m.Bytes.IsNegative() {
		return errors.Wrap(ErrorInvalidBytes, "bytes cannot be negative")
	}

	return nil
}

func (m *MsgUpdateQuotaRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
