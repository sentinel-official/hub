package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// The `types` package contains custom message types for the Cosmos SDK.

// The following variables implement the sdk.Msg interface for MsgCancelRequest and MsgAllocateRequest.
// These variables ensure that the corresponding types can be used as messages in the Cosmos SDK.
var (
	_ sdk.Msg = (*MsgCancelRequest)(nil)
	_ sdk.Msg = (*MsgAllocateRequest)(nil)
)

// NewMsgCancelRequest creates a new MsgCancelRequest instance with the given parameters.
func NewMsgCancelRequest(from sdk.AccAddress, id uint64) *MsgCancelRequest {
	return &MsgCancelRequest{
		From: from.String(),
		ID:   id,
	}
}

// ValidateBasic performs basic validation checks on the MsgCancelRequest fields.
// It checks if the 'From' field is not empty and represents a valid account address,
// and if the 'ID' field is not zero.
func (m *MsgCancelRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrapf(ErrorInvalidMessage, "invalid from %s", err)
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgCancelRequest.
func (m *MsgCancelRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

// NewMsgAllocateRequest creates a new MsgAllocateRequest instance with the given parameters.
func NewMsgAllocateRequest(from sdk.AccAddress, id uint64, addr sdk.AccAddress, bytes sdkmath.Int) *MsgAllocateRequest {
	return &MsgAllocateRequest{
		From:    from.String(),
		ID:      id,
		Address: addr.String(),
		Bytes:   bytes,
	}
}

// ValidateBasic performs basic validation checks on the MsgAllocateRequest fields.
// It checks if the 'From' field is not empty and represents a valid account address,
// if the 'ID' field is not zero,
// if the 'Address' field is not empty and represents a valid account address,
// and if the 'Bytes' field is not nil and not negative.
func (m *MsgAllocateRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrapf(ErrorInvalidMessage, "invalid from %s", err)
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Address == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return sdkerrors.Wrapf(ErrorInvalidMessage, "invalid address %s", err)
	}
	if m.Bytes.IsNil() {
		return sdkerrors.Wrap(ErrorInvalidMessage, "bytes cannot be nil")
	}
	if m.Bytes.IsNegative() {
		return sdkerrors.Wrap(ErrorInvalidMessage, "bytes cannot be negative")
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgAllocateRequest.
func (m *MsgAllocateRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
