package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgSubscribeToNodeRequest)(nil)
	_ sdk.Msg = (*MsgSubscribeToPlanRequest)(nil)
	_ sdk.Msg = (*MsgCancelRequest)(nil)

	_ sdk.Msg = (*MsgAddQuotaRequest)(nil)
	_ sdk.Msg = (*MsgUpdateQuotaRequest)(nil)
)

func NewMsgSubscribeToNodeRequest(from sdk.AccAddress, address hubtypes.NodeAddress, deposit sdk.Coin) *MsgSubscribeToNodeRequest {
	return &MsgSubscribeToNodeRequest{
		From:    from.String(),
		Address: address.String(),
		Deposit: deposit,
	}
}

func (m *MsgSubscribeToNodeRequest) Route() string {
	return RouterKey
}

func (m *MsgSubscribeToNodeRequest) Type() string {
	return TypeMsgSubscribeToNodeRequest
}

func (m *MsgSubscribeToNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidAddress, "address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidAddress, "%s", err)
	}
	if m.Deposit.IsNegative() {
		return errors.Wrap(ErrorInvalidDeposit, "deposit cannot be negative")
	}
	if m.Deposit.IsZero() {
		return errors.Wrap(ErrorInvalidDeposit, "deposit cannot be zero")
	}
	if !m.Deposit.IsValid() {
		return errors.Wrap(ErrorInvalidDeposit, "deposit must be valid")
	}

	return nil
}

func (m *MsgSubscribeToNodeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgSubscribeToNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgSubscribeToPlanRequest(from sdk.AccAddress, id uint64, denom string) *MsgSubscribeToPlanRequest {
	return &MsgSubscribeToPlanRequest{
		From:  from.String(),
		Id:    id,
		Denom: denom,
	}
}

func (m *MsgSubscribeToPlanRequest) Route() string {
	return RouterKey
}

func (m *MsgSubscribeToPlanRequest) Type() string {
	return TypeMsgSubscribeToPlanRequest
}

func (m *MsgSubscribeToPlanRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return errors.Wrapf(ErrorInvalidDenom, "%s", err)
		}
	}

	return nil
}

func (m *MsgSubscribeToPlanRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgSubscribeToPlanRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgCancelRequest(from sdk.AccAddress, id uint64) *MsgCancelRequest {
	return &MsgCancelRequest{
		From: from.String(),
		Id:   id,
	}
}

func (m *MsgCancelRequest) Route() string {
	return RouterKey
}

func (m *MsgCancelRequest) Type() string {
	return TypeMsgCancelRequest
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

func (m *MsgCancelRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgCancelRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgAddQuotaRequest(from sdk.AccAddress, id uint64, address sdk.AccAddress, bytes sdk.Int) *MsgAddQuotaRequest {
	return &MsgAddQuotaRequest{
		From:    from.String(),
		Id:      id,
		Address: address.String(),
		Bytes:   bytes,
	}
}

func (m *MsgAddQuotaRequest) Route() string {
	return RouterKey
}

func (m *MsgAddQuotaRequest) Type() string {
	return TypeMsgAddQuotaRequest
}

func (m *MsgAddQuotaRequest) ValidateBasic() error {
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

func (m *MsgAddQuotaRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgAddQuotaRequest) GetSigners() []sdk.AccAddress {
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

func (m *MsgUpdateQuotaRequest) Route() string {
	return RouterKey
}

func (m *MsgUpdateQuotaRequest) Type() string {
	return TypeMsgUpdateQuotaRequest
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

func (m *MsgUpdateQuotaRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateQuotaRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
